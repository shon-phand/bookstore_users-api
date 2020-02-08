package users

import (
	"fmt"

	"github.com/shon-phand/bookstore_users-api/dataSources/mysql/users_db"
	"github.com/shon-phand/bookstore_users-api/domains/errors"
)

const (
	queryInsertUser = "INSERT INTO users (first_name,last_name,email,date_created) VALUES( ?,?,?,? )"
)

var (
	UsersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

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
	fmt.Println("saving user started")

	conn := users_db.Client
	err := conn.Ping()
	if err != nil {
		panic(err)
	}
	// stmt, err := users_db.Client.Prepare(queryInsertUser)
	// fmt.Println("stmt:", stmt,"err:",err)
	// if err != nil {
	// 	//return errors.StatusInternalServerError("error in preapre stmt : " + err.Error())
	// }

	// defer stmt.Close()
	//user.CreationDate = date_utils.GetNowString()
	//users_db.Client.Ping()

	// err := users_db.Client.Ping()
	// fmt.Println("error after pinged", err)
	// // if err != nil {
	// 	panic(err)
	// }
	fmt.Println("database pinged")

	// fmt.Println("database pinged")
	// 	insertResult, err := users_db.Client.Exec(queryInsertUser,user.FirstName, user.LastName, user.Email, user.CreationDate)
	// 	if err != nil {
	// 		return errors.StatusInternalServerError("error while saving user : " + err.Error())
	// 	}

	// 	userId, err := insertResult.LastInsertId()
	// 	if err != nil {
	// 		return errors.StatusInternalServerError("error while saving record" + err.Error())
	// 	}
	// 	user.ID = userId
	// 	return nil

	// current := UsersDB[user.ID]
	// if current != nil {
	// 	if user.Email == current.Email {
	// 		return errors.StatusBadRequestError("Email  is already registered")
	// 	}
	// 	return errors.StatusBadRequestError("User-ID already exist")
	// }
	// user.CreationDate = date_utils.GetNowString()
	// UsersDB[user.ID] = user
	return nil
}
