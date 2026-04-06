package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RBACMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		roleVal, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		role, ok := roleVal.(string)
		if !ok || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "forbidden: admin only",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
