package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shon-phand/bookstore_users-api/controllers/configuration"
	"github.com/shon-phand/bookstore_users-api/domains/errors"
	"github.com/shon-phand/bookstore_users-api/logger"
)

var (
	r = gin.Default()
)

func StartApp() {
	fmt.Println("starting webserver")
	config, err := configuration.LoadConfiguration("/home/shon/Documents/Microservice/golang-microservice/src/github.com/shon-phand/bookstore_users-api/app/webserver_properties.json")
	if err != nil {
		logger.Info(errors.StatusInternalServerError("unable to load configuration file"), err)
		panic(err)
	}
	mapUrls()
	r.Run(config.Host + ":" + config.Port)
}
