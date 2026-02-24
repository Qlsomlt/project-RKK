package middleware

import (
	"strings"

	"kode/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Error(c, 401, "authorization header missing")
			c.Abort()
			return
		}

		// Expect: Bearer <token>
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			utils.Error(c, 401, "invalid authorization format")
			c.Abort()
			return
		}

		tokenString := splitToken[1]

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			utils.Error(c, 401, "invalid or expired token")
			c.Abort()
			return
		}

		// Store user data in context
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
