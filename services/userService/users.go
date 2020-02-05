package userService

import (
	//"fmt"

	"github.com/shon-phand/bookstore_users-api/domains/errors"
	"github.com/shon-phand/bookstore_users-api/domains/users"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {

		result:= &users.User{ID: userId}

		if err := result.Get(); err!= nil{
			return nil, err
		}
	

	return result, nil

}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	//fmt.Println("validating email")
	if err := user.ValidateEmail(); err != nil {
		return nil, err
	}
	//fmt.Println("email validated")
	//fmt.Println("saving user")
	if err := user.Save(); err != nil {
		return nil, err
	}
	//fmt.Println("user saved")
	return &user, nil

}
