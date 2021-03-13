package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/shon-phand/bookstore_users-api/domains/errors"
	"github.com/shon-phand/bookstore_users-api/logger"
)

var (
	Client *sql.DB
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "shon1234"
	dbname   = "postgres"
)

func init() {

	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	Client, err = sql.Open("postgres", dataSourceName)
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
