package users

import (
	"strings"

	"github.com/shon-phand/bookstore_users-api/domains/errors"
	"github.com/shon-phand/bookstore_users-api/logger"
)

type User struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	CreationDate string `json:"date_created"`
	Password     string `json:"password"`
	Status       string `json:"status"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.Password = strings.TrimSpace((user.Password))
	if user.Email == "" {
		logger.Info(errors.StatusBadRequestError("Invalid email address"), nil)
		return errors.StatusBadRequestError("Invalid email address")
	}
	if user.Password == "" {
		logger.Info(errors.StatusBadRequestError("Invalid password"), nil)
		return errors.StatusBadRequestError("Invalid password")
	}
	return nil

}
