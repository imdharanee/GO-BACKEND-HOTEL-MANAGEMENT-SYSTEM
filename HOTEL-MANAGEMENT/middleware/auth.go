package middleware

import (
	"HOTEL-MANAGEMENT/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clienttoken := c.Request.Header.Get("token")

		if clienttoken == " " {

			c.JSON(http.StatusInternalServerError, gin.H{"error": "No authentication required"})
			c.Abort()
			return

		}
		claims, err := helpers.ValidateTokens(clienttoken)

		if err != " " {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("name", claims.Name)

		c.Set("uid", claims.Uid)

	}
}
