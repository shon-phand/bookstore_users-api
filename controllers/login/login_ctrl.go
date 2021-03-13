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

var expirationTime = time.Now().Add(5 * time.Minute)

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
		tokenString, err := CreateToken(creds.Username)
		if err != nil {
			c.JSON(500, "error in crating token")
			return
		}
		c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
		//fmt.Println(c.Writer)
		c.JSON(200, "login successful")
	}
}

func CreateToken(username string) (string, error) {

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: username,
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
		//c.JSON(500, "error in generating jwt authentication token")
		return "", err
	}
	return tokenString, nil

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

func RefreshToken(c *gin.Context) {

	isValid := ValidateToken(c)
	if !isValid {
		fmt.Println("Token is not valid")
		return
	}
	// (END) The code up-till this point is the same as the first part of the `Welcome` route
	claims := &Claims{}
	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		//w.WriteHeader(http.StatusBadRequest)
		//return "", err
		fmt.Println("Yet to expire, more than 30 sec")
		return
	}
	// Now, create a new token for the current use, with a renewed expiration time
	//expirationTime := time.Now().Add(1 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		//w.WriteHeader(http.StatusInternalServerError)
		//return "", err
		fmt.Println("500")
		return
	}

	// Set the new token as the users `token` cookie
	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
	// })
	//fmt.Println("token refreshed")
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
}
