package middleware

import (
	"kode/utils"

	"github.com/gin-gonic/gin"
)

func Authorize(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		role := c.GetString("role")
		if role == "" {
			utils.Error(c, 403, "forbidden")
			c.Abort()
			return
		}

		for _, allowedRole := range roles {
			if role == allowedRole {
				c.Next()
				return
			}
		}

		utils.Error(c, 403, "access denied")
		c.Abort()
	}
}
