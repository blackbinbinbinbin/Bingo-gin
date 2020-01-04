package router

import (
	"github.com/gin-gonic/gin"
	"Bingo-gin/router/middleware/exception"
	"Bingo-gin/router/middleware/logger"
)

type Router struct {
	Root           *gin.Engine
	IndexApi          *gin.RouterGroup		//首页Api
}

var BaseRouter *Router

func InitRouter(router *gin.Engine) {
	//设置路由中间件
	router.Use(exception.SetUp(), logger.SetUp())

	BaseRouter := &Router{Root: router}
	//首页Api
	BaseRouter.Index()

	return
}
