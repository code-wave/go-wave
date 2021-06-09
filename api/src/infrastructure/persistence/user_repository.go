package persistence

import (
	"database/sql"
	"log"
	"strings"

	"github.com/code-wave/go-wave/domain/entity"
	"github.com/code-wave/go-wave/domain/repository"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

const (
	querySaveUser = "INSERT INTO users (email, password, name, nickname, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id;"
)

var _ repository.UserRepository = &UserRepo{}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Save(user *entity.User) (*entity.User, *errors.RestErr) {
	stmt, err := r.db.Prepare(querySaveUser)
	if err != nil {
		log.Println("error when trying to prepare to save user", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if err = stmt.QueryRow(user.Email, user.Password, user.Name, user.Nickname, user.CreatedAt).
		Scan(&user.ID); err != nil {
		log.Println("error when trying to scan to save user", err)
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.NewBadRequestError("email is duplicated, already taken")
		}
		return nil, errors.NewInternalServerError("database error")
	}

	return user, nil
}
