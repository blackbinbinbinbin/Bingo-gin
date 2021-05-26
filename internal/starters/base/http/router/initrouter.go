package router

import (
	"github.com/gin-gonic/gin"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters/base/http/middleware/exception"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters/base/http/middleware/logger"
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
