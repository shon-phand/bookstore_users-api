package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {

	dataSoueceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"root",
		"mypass",
		"127.0.0.1:3306",
		"bookstore_users",
	)
	Client, err := sql.Open("mysql", dataSoueceName)
	if err != nil {
		panic(err)
	}
	err = Client.Ping()
	if err != nil {
		fmt.Println("ping error")
		panic(err)
	}

	log.Println("Database successully connected")
}
