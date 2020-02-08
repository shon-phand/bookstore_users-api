package users

import (
	"strings"

	"github.com/shon-phand/bookstore_users-api/domains/errors"
)

type User struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	CreationDate string `json:"date_created"`
}

func (user *User) ValidateEmail() *errors.RestErr {

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.StatusBadRequestError("Invalid email address")
	}

	return nil

}
