package app

import (
	"github.com/shon-phand/bookstore_users-api/controllers/login"
	"github.com/shon-phand/bookstore_users-api/controllers/user"
)

func mapUrls() {
	r.GET("/ping", user.Ping())

	r.GET("/users/:user_id", user.GetUser())
	r.POST("/users", user.CreateUser())
	r.PUT("/users/:user_id", user.UpdateUser())
	r.PATCH("/users/:user_id", user.UpdateUser())
	r.DELETE("/users/:user_id", user.DeleteUser())
	r.GET("/internal/users", user.Search())
	r.POST("/login", login.Login())
}
