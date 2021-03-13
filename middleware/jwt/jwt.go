package jwt

import (
	"net/http"

	"github.com/shon-phand/bookstore_users-api/utils/jwt"

	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		isValid := jwt.ValidateToken(c)
		if !isValid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		jwt.RefreshToken(c)
	}
}
