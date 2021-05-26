package router

import (
	"github.com/blackbinbinbinbin/Bingo-gin/apis/http/my-app/v1"
)

func (r *Router) Index() {
	// 设置路由分组
	r.IndexApi = r.Root.Group("")
	// 设置controller
	//controller := controller.IndexController{}


	r.IndexApi.GET("/index", v1.IndexAction)
}
