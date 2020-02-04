package app

import "github.com/shon-phand/bookstore_users-api/controllers"

func mapUrls(){
	r.GET("/ping",controllers.Ping)
}
