package entity

import (
	"database/sql"
	"encoding/json"
	"strings"
	"time"

	"github.com/code-wave/go-wave/infrastructure/encryption"
	"github.com/code-wave/go-wave/infrastructure/errors"
	"github.com/code-wave/go-wave/infrastructure/helpers"
)

type Users []User

type User struct {
	ID        int64          `json:"id"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Name      string         `json:"name"`
	Nickname  string         `json:"nickname"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(user.Email)
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Password = strings.TrimSpace(user.Password)

	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address, email is required")
	}

	if user.Password == "" {
		return errors.NewBadRequestError("invalid password, password is required")
	} else {
		if len(user.Password) < 6 {
			return errors.NewBadRequestError("password should be at least 6 characters")
		}
	}

	return nil
}

func (u *User) BeforeSave() *errors.RestErr {
	u.Password, _ = encryption.Hash(u.Password)
	u.CreatedAt = helpers.GetDateString(time.Now())

	return nil
}

func (users Users) PublicUsers() []interface{} {
	result := make([]interface{}, len(users))
	for idx, user := range users {
		result[idx] = user.PublicUser()
	}
	return result
}

func (u *User) PublicUser() interface{} {
	return &PublicUser{
		ID:       u.ID,
		Email:    u.Email,
		Nickname: u.Nickname,
	}
}

func (u *User) ResponseJSON() interface{} {
	user := u.PublicUser()
	uJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return uJSON
}

func (u *Users) ResponseJSON() interface{} {
	users := u.PublicUsers()
	uJSON, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return uJSON
}
