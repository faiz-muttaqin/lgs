package middleware

import (
	"github.com/gin-gonic/gin"
)

func RouteGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
