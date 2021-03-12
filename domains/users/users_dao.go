package users

import (
	"fmt"
	"strings"

	"github.com/shon-phand/bookstore_users-api/dataSources/mysql/users_db"
	"github.com/shon-phand/bookstore_users-api/domains/errors"
	"github.com/shon-phand/bookstore_users-api/logger"
	"github.com/shon-phand/bookstore_users-api/utils/date_utils"
)

const (
	queryInsertUser     = "INSERT INTO users (first_name,last_name,email,password,status,date_created) VALUES( $1,$2,$3,$4,$5,$6 )"
	queryGetUser        = "SELECT id,first_name,last_name,email,password,status,date_created FROM users where id=$1 "
	queryUpdateUser     = "UPDATE  users SET first_name=$1,last_name=$2,email=$3 WHERE id=$4;"
	queryDeleteUser     = "DELETE from users WHERE id=$1;"
	queryFindByStatus   = "SELECT id,first_name,last_name,email,status,date_created FROM users WHERE status=$1;"
	queryGetUserByEmail = "SELECT id,first_name,last_name,email,password,status,date_created FROM users where email=$1 "
)

func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUser)
	//fmt.Println("stmt:", stmt, "err:", err)
	if err != nil {
		logger.Error(errors.StatusInternalServerError("database error"), err)
		return errors.StatusInternalServerError("database error ")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status, &user.CreationDate); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			logger.Info(errors.StatusNotFoundError("users-id not found"), err)
			return errors.StatusNotFoundError("users-id not found")
		}
		logger.Error(errors.StatusInternalServerError("error in fetching data"), err)
		return errors.StatusInternalServerError("database error")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	//fmt.Println("stmt:", stmt, "err:", err)
	if err != nil {
		logger.Error(errors.StatusInternalServerError("error in preapre stmt"), err)
		return errors.StatusInternalServerError(" database error")
	}

	defer stmt.Close()
	user.CreationDate = date_utils.GetNowString()
	//fmt.Println("user creating : ", user)
	_, err = stmt.Exec(&user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status, &user.CreationDate)
	//insertResult, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.CreationDate)
	if err != nil {
		return errors.StatusInternalServerError("error while saving user : " + err.Error())
	}
	//fmt.Println(insertResult)
	// userId, err := insertResult.LastInsertId()
	// if err != nil {
	// 	return errors.StatusInternalServerError("error while saving record" + err.Error())
	// }

	stmt, err = users_db.Client.Prepare(queryGetUserByEmail)
	//fmt.Println("stmt:", stmt, "err:", err)
	if err != nil {
		logger.Error(errors.StatusInternalServerError("database error"), err)
		return errors.StatusInternalServerError("database error ")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status, &user.CreationDate); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			logger.Info(errors.StatusNotFoundError("email not found"), err)
			return errors.StatusNotFoundError("email not found")
		}
		logger.Error(errors.StatusInternalServerError("error in fetching data"), err)
		return errors.StatusInternalServerError("database error")
	}

	//user.ID = userId
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
		return errors.StatusInternalServerError("error in deleting user")
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		return nil, errors.StatusInternalServerError("error in preapre stmt : " + err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.StatusInternalServerError("error in fetching data")
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.CreationDate)
		if err != nil {
			errors.StatusInternalServerError("error in scanning data")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.StatusNotFoundError("no users with given status")
	}

	return results, nil

}

func GetUserByUsername(email string) (*User, *errors.RestErr) {

	stmt, err := users_db.Client.Prepare(queryGetUserByEmail)
	//fmt.Println("stmt:", stmt, "err:", err)
	if err != nil {
		logger.Error(errors.StatusInternalServerError("database error"), err)
		return nil, errors.StatusInternalServerError("database error ")
	}
	defer stmt.Close()
	var user User
	fmt.Println("searching for email : ", email)
	result := stmt.QueryRow(email)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status, &user.CreationDate); err != nil {
		logger.Info(errors.StatusNotFoundError("email not found"), err)
		return nil, errors.StatusNotFoundError("email not found")

	}
	fmt.Println(user)
	return &user, nil

}
