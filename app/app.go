package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	r = gin.Default()
)

func StartApp() {
	fmt.Println("starting webserver")
	mapUrls()
	err := r.Run(":8000")

	if err != nil {
		panic("HTTP server failed to start")
	}
}
