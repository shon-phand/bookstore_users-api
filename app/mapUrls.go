package app

import (
	"github.com/shon-phand/bookstore_users-api/controllers/login"
	"github.com/shon-phand/bookstore_users-api/controllers/user"
	"github.com/shon-phand/bookstore_users-api/middleware/jwt"
)

func mapUrls() {
	r.GET("/ping", user.Ping())

	// route which create create JWT token
	r.POST("/login", login.Login())

	// require JWT authentication for below routes
	authorized := r.Group("/")
	authorized.Use(jwt.Jwt())
	{
		authorized.GET("/users/:user_id", user.GetUser())
		authorized.PUT("/users/:user_id", user.UpdateUser())
		authorized.PATCH("/users/:user_id", user.UpdateUser())
		authorized.DELETE("/users/:user_id", user.DeleteUser())
		authorized.GET("/users", user.Search())
	}
	// JWT authentication not required for below routes
	publicRoutes := r.Group("/")
	{
		publicRoutes.POST("/users", user.CreateUser())
	}

}
