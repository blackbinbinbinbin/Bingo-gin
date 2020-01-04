package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Bingo-gin/common"
)

type BaseController struct{
	context *gin.Context
}

func (ctl BaseController) buildSuccessData(msg string, data interface{}) map[string]interface{} {
	return gin.H{
		"code":   common.CodeSuccess,
		"msg" :   msg,
		"data":   data,
	}
}

func (ctl BaseController) buildErrorData(code int, msg string, data interface{}) map[string]interface{} {
	return gin.H{
		"code" :  code,
		"msg"  :  msg,
		"data" :  data,
	}
}

func (ctl BaseController) SuccessResponse(message string, data interface{}) {
	ctl.context.JSON(http.StatusOK, ctl.buildSuccessData(message, data))
}

func (ctl BaseController) ErrorResponse(statusCode int, message string, data interface{}) {
	ctl.context.JSON(http.StatusOK, ctl.buildErrorData(statusCode, message, data))
}