package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"error": "authorization header required",
			})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)
		claims, err := ValidateToken(tokenString)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
