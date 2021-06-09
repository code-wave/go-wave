package entity

import (
	"database/sql"
	"encoding/json"
	"strings"
	"time"

	"github.com/code-wave/go-wave/infrastructure/date"
	"github.com/code-wave/go-wave/infrastructure/encryption"
	"github.com/code-wave/go-wave/infrastructure/errors"
)

type Users []User

type User struct {
	ID        uint64         `json:"id"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Name      string         `json:"name"`
	Nickname  string         `json:"nickname"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
}

type PublicUser struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
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
	u.CreatedAt = date.GetDateString(time.Now())
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
		Name:     u.Name,
		Nickname: u.Nickname,
	}
}

func (u *User) ResponseJSON() interface{} {
	user := u.PublicUser()
	uJSON, err := json.Marshal(user)
	if err != nil {
		return nil
	}
	return uJSON
}
