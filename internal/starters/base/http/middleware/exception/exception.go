package exception

import "github.com/gin-gonic/gin"

func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}