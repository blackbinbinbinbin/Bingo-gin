package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexAction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code" :  200,
		"msg"  :  "success",
		"data" :  nil,})
	return
}