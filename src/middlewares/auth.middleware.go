package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"e-shop-backend/src/shared"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token", "status": false})
			c.Abort()
			return
		}

		claims, err := shared.ValidateToken(tokenString)

		if shared.HandleError(c, err) {
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
