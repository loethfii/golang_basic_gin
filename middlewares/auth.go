package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang_basic_gin/auth"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request Access Token"})
			c.Abort()
			return
		}
		//validte token

		_, _, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "UnAuthorized",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
