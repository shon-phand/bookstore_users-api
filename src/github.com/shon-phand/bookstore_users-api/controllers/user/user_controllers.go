package user

import (
	//"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shon-phand/bookstore_users-api/domains/errors"
	"github.com/shon-phand/bookstore_users-api/domains/users"
	"github.com/shon-phand/bookstore_users-api/services/userService"
)

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "pong")
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if userErr != nil {
			err := errors.StatusBadRequestError("userID should be number value")
			c.JSON(err.Status, err)
			return
		}

		user, getErr := userService.GetUser(userId)
		if getErr != nil {
			c.JSON(getErr.Status, getErr)
			return
		}

		c.JSON(http.StatusOK, user.Marshall(c.GetHeader("x-public") == "true"))
	}
}

func CreateUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		var user users.User
		err := c.ShouldBindJSON(&user)
		//fmt.Println("json binded")
		if err != nil {

			resterr := errors.StatusBadRequestError("invalid request json")

			c.JSON(resterr.Status, resterr.Message)
			return

		}
		//fmt.Println("callinng CreateUser service")
		result, saveErr := userService.CreateUser(user)
		if saveErr != nil {

			c.JSON(saveErr.Status, saveErr)
			return
		}

		c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("x-public") == "true"))
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if userErr != nil {
			err := errors.StatusBadRequestError("userID should be number value")
			c.JSON(err.Status, err)
			return
		}

		var user users.User
		err := c.ShouldBindJSON(&user)
		//fmt.Println("json binded")
		if err != nil {

			resterr := errors.StatusBadRequestError("invalid request json")
			c.JSON(resterr.Status, resterr.Message)
			return
		}
		user.ID = userId

		isPartial := c.Request.Method == http.MethodPatch

		result, updateErr := userService.UpdateUser(isPartial, user)
		if updateErr != nil {
			c.JSON(updateErr.Status, updateErr)
			return
		}
		c.JSON(http.StatusOK, result.Marshall(c.GetHeader("x-public") == "true"))
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if userErr != nil {
			err := errors.StatusBadRequestError("userID should be number value")
			c.JSON(err.Status, err)
			return
		}

		var user users.User
		err := c.ShouldBindJSON(&user)
		//fmt.Println("json binded")
		if err != nil {

			resterr := errors.StatusBadRequestError("invalid request json")
			c.JSON(resterr.Status, resterr.Message)
			return
		}
		user.ID = userId

		_, deleteErr := userService.DeleteUser(user)
		if deleteErr != nil {
			c.JSON(deleteErr.Status, deleteErr)
			return
		}
		c.JSON(http.StatusOK, map[string]string{
			"message": "user deleted",
		})
	}
}

func Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Query("status")

		data, err := userService.Search(status)

		if err != nil {
			c.JSON(err.Status, err)
			return
		}
		c.JSON(http.StatusOK, data.Marshall(c.GetHeader("x-public") == "true"))
	}
}
