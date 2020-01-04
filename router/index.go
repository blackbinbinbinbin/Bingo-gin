package router

import (
	"Bingo-gin/controller"
)

func (r *Router) Index() {
	// 设置路由分组
	r.IndexApi = r.Root.Group("")
	// 设置controller
	controller := controller.IndexController{}


	r.IndexApi.GET("/index", controller.IndexAction)
}