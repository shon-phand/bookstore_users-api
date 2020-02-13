package users

import (
	"fmt"
	"strings"

	"github.com/shon-phand/bookstore_users-api/dataSources/mysql/users_db"
	"github.com/shon-phand/bookstore_users-api/domains/errors"
	"github.com/shon-phand/bookstore_users-api/utils/date_utils"
)

const (
	queryInsertUser = "INSERT INTO users (first_name,last_name,email,date_created) VALUES( ?,?,?,? )"
	queryGetUser    = "SELECT id,first_name,last_name,email,date_created FROM users where id=? "
	queryUpdateUser = "UPDATE  users SET first_name=?,last_name=?,email=? WHERE id=?;"
	queryDeleteUser = "DELETE from users WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUser)
	//fmt.Println("stmt:", stmt, "err:", err)
	if err != nil {
		return errors.StatusInternalServerError("error in preapre stmt : " + err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreationDate); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.StatusNotFoundError("users-id not found")
		}
		return errors.StatusInternalServerError("error in fetching data")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	fmt.Println("stmt:", stmt, "err:", err)
	if err != nil {
		return errors.StatusInternalServerError("error in preapre stmt : " + err.Error())
	}

	defer stmt.Close()
	user.CreationDate = date_utils.GetNowString()

	insertResult, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.CreationDate)
	if err != nil {
		return errors.StatusInternalServerError("error while saving user : " + err.Error())
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.StatusInternalServerError("error while saving record" + err.Error())
	}
	user.ID = userId
	return nil

}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.StatusInternalServerError("error in preapre stmt : " + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(&user.FirstName, &user.LastName, &user.Email, &user.ID)
	if err != nil {
		return errors.StatusInternalServerError("error in updating user")
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.StatusInternalServerError("error in preapre stmt : " + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(&user.ID)
	if err != nil {
		return errors.StatusInternalServerError("error in updating user")
	}

	return nil
}
