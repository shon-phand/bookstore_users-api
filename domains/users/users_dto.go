package users

import (
	"github.com/shon-phand/bookstore_users-api/domains/errors"
	"github.com/shon-phand/bookstore_users-api/utils/date_utils"
)
var (
	UsersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := UsersDB[user.ID]
	if result != nil {
		user.ID = result.ID
		user.FirstName = result.FirstName
		user.LastName = result.LastName
		user.Email = result.Email

		return nil
	}
	return errors.StatusNotFoundError("user-id not found")
}

func (user *User) Save() *errors.RestErr {
	current := UsersDB[user.ID]
	if current != nil {
		if user.Email == current.Email {
			return errors.StatusBadRequestError("Email  is already registered")
		}
		return errors.StatusBadRequestError("User-ID already exist")
	}
	user.CreationDate = date_utils.GetNowString()
	UsersDB[user.ID] = user
	return nil
}
