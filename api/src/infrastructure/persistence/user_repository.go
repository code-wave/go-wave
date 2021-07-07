package persistence

import (
	"database/sql"
	"log"
	"strings"

	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/encryption"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"golang.org/x/crypto/bcrypt"
)

const (
	querySaveUser               = "INSERT INTO users (email, password, name, nickname, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id;"
	queryGetUserByID            = "SELECT id, email, name, nickname, created_at, updated_at FROM users WHERE id = $1;"
	queryGetAllUsers            = "SELECT * FROM users LIMIT $1 OFFSET $2;"
	queryFindByEmailAndPassword = "SELECT id, email, password, name, nickname, created_at, updated_at FROM users WHERE email = $1;"
	queryUpdateUser             = "UPDATE users SET password = $1, name = $2, nickname = $3, updated_at = $4 WHERE id = $5;"
	queryDeleteUser             = "DELETE FROM users WHERE id = $1;"
)

var _ repository.UserRepository = &UserRepo{}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Save(user *entity.User) *errors.RestErr {
	stmt, err := r.db.Prepare(querySaveUser)
	if err != nil {
		log.Println("error when trying to prepare to save user, ", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if err = stmt.QueryRow(user.Email, user.Password, user.Name, user.Nickname, user.CreatedAt).
		Scan(&user.ID); err != nil {
		log.Println("error when trying to scan to save user, ", err)
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return errors.NewBadRequestError("email is duplicated, already taken")
		}
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (r *UserRepo) GetUserByID(userID int64) (*entity.User, *errors.RestErr) {
	stmt, err := r.db.Prepare(queryGetUserByID)
	if err != nil {
		log.Println("error when trying to prepare to get user by id, ", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	user := entity.User{
		ID: userID,
	}

	if err = stmt.QueryRow(user.ID).Scan(&user.ID, &user.Email, &user.Name, &user.Nickname, &user.CreatedAt, &user.UpdatedAt); err != nil {
		log.Println("error when trying to scan after get user by id, ", err)
		return nil, errors.NewInternalServerError("database error")
	}
	return &user, nil
}

func (r *UserRepo) Get(user *entity.User) *errors.RestErr {
	stmt, err := r.db.Prepare(queryGetUserByID)
	if err != nil {
		log.Println("error when trying to prepare to get user by id, ", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if err = stmt.QueryRow(user.ID).Scan(&user.ID, &user.Email, &user.Name, &user.Nickname, &user.CreatedAt, &user.UpdatedAt); err != nil {
		log.Println("error when trying to scan after get user by id, ", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (r *UserRepo) GetAll(limit, offset int64) (entity.Users, *errors.RestErr) {
	stmt, err := r.db.Prepare(queryGetAllUsers)
	if err != nil {
		log.Println("error when trying to prepare to get all users with limit & offset, ", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		log.Println("error when trying to Query to get all users with limit & offset, ", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	users := make(entity.Users, 0)
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Nickname, &user.CreatedAt, &user.UpdatedAt); err != nil {
			log.Println("error when trying to scan to get all users, ", err)
			return nil, errors.NewInternalServerError("database error")
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepo) Update(user *entity.User) *errors.RestErr {
	stmt, err := r.db.Prepare(queryUpdateUser)
	if err != nil {
		log.Println("error when trying to prepare to update user, ", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if user.UpdatedAt.Valid {
		_, err = stmt.Exec(user.Password, user.Name, user.Nickname, user.UpdatedAt.String, user.ID)
	} else {
		_, err = stmt.Exec(user.Password, user.Name, user.Nickname, nil, user.ID)
	}

	if err != nil {
		log.Println("error when trying to execute update user, ", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (r *UserRepo) Delete(userID int64) *errors.RestErr {
	stmt, err := r.db.Prepare(queryDeleteUser)
	if err != nil {
		log.Println("error when trying to prepare to delete user, ", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		log.Println("error when trying to exeucte delete user, ", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (r *UserRepo) FindByEmailAndPassword(lu *entity.User) (*entity.User, *errors.RestErr) {
	stmt, err := r.db.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		log.Println("error when trying to prepare to find user by email and password, ", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	var user entity.User
	if err := stmt.QueryRow(lu.Email).
		Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Nickname, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, errors.NewNotFoundError("does not exist that email in database")
		}
		log.Println("error when trying to find user by email and password after scan, ", err)
		return nil, errors.NewInternalServerError("database error")
	}

	if err := encryption.VerifyPassword(user.Password, lu.Password); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, errors.NewNotFoundError("given password does not match database's password")
		}
		return nil, errors.NewInternalServerError("hashing password error")
	}

	return &user, nil
}
