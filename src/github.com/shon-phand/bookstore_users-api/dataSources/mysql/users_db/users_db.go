package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shon-phand/bookstore_users-api/domains/errors"
	"github.com/shon-phand/bookstore_users-api/logger"
)

var (
	Client *sql.DB
)

func init() {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"root",
		"mysqlpasswd",
		"127.0.0.1:3306",
		"bookstore_users",
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		logger.Error(errors.StatusInternalServerError("error in starting database"), err)
		panic(err)
	}
	err = Client.Ping()
	if err != nil {
		logger.Error(errors.StatusInternalServerError("error in pinging database"), err)
		panic(err)
	}

	log.Println("Database successully connected")
}
