package login

import (
	"fmt"
	

	"github.com/gin-gonic/gin"
	"github.com/shon-phand/bookstore_users-api/domains/errors"
	"github.com/shon-phand/bookstore_users-api/domains/users"
	"github.com/shon-phand/bookstore_users-api/logger"
	"github.com/shon-phand/bookstore_users-api/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login() gin.HandlerFunc {

	return func(c *gin.Context) {
		var creds jwt.Credentials
		err := c.ShouldBindJSON(&creds)
		if err != nil {
			resterr := errors.StatusBadRequestError("invalid request json")
			logger.Info(resterr, nil)
			c.JSON(resterr.Status, resterr.Message)
			return
		}

		var user *users.User

		user, userErr := users.GetUserByUsername(creds.Username)
		//fmt.Println("fetched user", user)
		if userErr != nil {
			//fmt.Println("unable  to fetch user details in login", userErr)
			//logger.Info(userErr.Status, nil)
			c.JSON(400, "login failed")
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
		if err != nil {
			fmt.Println("password not match")
			c.JSON(400, "password not match")
			return
		}
		tokenString, err := jwt.CreateToken(creds.Username)
		if err != nil {
			c.JSON(500, "error in crating token")
			return
		}
		c.SetCookie("token", tokenString, 300, "/", "localhost", false, true)
		//fmt.Println(c.Writer)
		data:= make(map[string]string)
		data["token"]=tokenString
		data["user"]=creds.Username
		c.JSON(200, data)
	}
}

