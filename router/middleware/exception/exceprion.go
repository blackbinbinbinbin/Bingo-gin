package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"Bingo-gin/util/response"
	"log"
	"Bingo-gin/common"
)

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				// 处理异常
				log.Println("【Error】系统错误：")
				log.Println(fmt.Sprintf("ErrorMsg：%s", err))
				log.Println(fmt.Sprintf("RequestURL：%s  %s%s", c.Request.Method, c.Request.Host, c.Request.RequestURI))
				log.Println(fmt.Sprintf("RequestUA：%s", c.Request.UserAgent()))
				log.Println("DebugStack：")
				log.Println(string(debug.Stack()))

				utilGin := response.Gin{Ctx: c}
				errMsg := "系统异常"
				cfg := common.NewConfig()
				appMode := cfg.GetValue("", "APP_ENV_MODE")
				// 如果打开测试开关，将详细信息返回至响应中
				if appMode == "debug" {
					errMsg = errMsg + fmt.Sprintf("：%s", err)
				}

				code := common.CodeSystemError
				switch t:=err.(type) {
				case *common.Error:
					code = t.Code
				}
				utilGin.Response(code, errMsg, nil)
			}
		}()
		c.Next()
	}
}