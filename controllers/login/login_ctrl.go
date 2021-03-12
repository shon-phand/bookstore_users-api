package login

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shon-phand/bookstore_users-api/domains/errors"
	"github.com/shon-phand/bookstore_users-api/domains/users"
	"github.com/shon-phand/bookstore_users-api/logger"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login() gin.HandlerFunc {

	return func(c *gin.Context) {
		var creds Credentials
		err := c.ShouldBindJSON(&creds)
		if err != nil {
			resterr := errors.StatusBadRequestError("invalid request json")
			logger.Info(resterr, nil)
			c.JSON(resterr.Status, resterr.Message)
			return
		}
		var user *users.User

		user, userErr := users.GetUserByUsername(creds.Username)
		fmt.Println("fetched user", user)
		if userErr != nil {
			fmt.Println("unable  to fetch user details in login", userErr)
			//logger.Info(userErr.Status, nil)
			c.JSON(400, "login failed")
			return
		}
		fmt.Println()
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
		if err != nil {
			fmt.Println("password not match")
			c.JSON(400, "password not match")
			return
		}

		expirationTime := time.Now().Add(2 * time.Minute)
		// Create the JWT claims, which includes the username and expiry time
		claims := &Claims{
			Username: creds.Username,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Create the JWT string
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			c.JSON(500, "error in generating jwt authentication token")
			return
		}
		fmt.Println("jwt ", tokenString)
		c.SetCookie("token", tokenString, int(claims.ExpiresAt), "/", "localhost", false, true)
		fmt.Println(c.Writer)
		c.JSON(200, "login successful")
	}

}

func ValidateToken(c *gin.Context) bool {
	cookie, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return false
		}
		// For any other type of error, return a bad request status
		return false
	}
	tknStr := cookie

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false
		}
		return false
	}
	if !tkn.Valid {
		return false
	}

	return true

}
