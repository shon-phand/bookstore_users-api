package userService

import (
	//"fmt"

	"github.com/shon-phand/bookstore_users-api/domains/errors"
	"github.com/shon-phand/bookstore_users-api/domains/users"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil

}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {

	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.Email = user.Email
		current.FirstName = user.FirstName
		current.LastName = user.LastName
	}

	if err = current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}
