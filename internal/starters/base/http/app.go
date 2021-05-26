package http

import (
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters/base/log"
	"os"
	"time"
	log2 "log"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters/base/http/router"
	"context"
)


var ginApplication *gin.Engine
var stopChan = make(chan bool)

func GinEngine() *gin.Engine {
	starters.Check(ginApplication)
	return ginApplication
}


type GinServerStart struct {
	starters.BaseStarter
}

func (g *GinServerStart) Init (ctx starters.StarterContext) {
}



func (g *GinServerStart) StartBlocking() bool {
	return true
}

func (g *GinServerStart) Setup(ctx starters.StarterContext) {
	logHelp := log.LoggerHelp()
	logHelp.Info("Gin Server Init()....")


	cfg := ctx.Props()
	env := cfg.GetDefault("app_env_mode", "develop")

	// 设置 gin 框架模式
	if env != "develop" {
		gin.SetMode(gin.ReleaseMode)
	}
	ginApplication := gin.New()
	ginApplication.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Bingo-gin")
	})

	// 初始化路由
	router.InitRouter(ginApplication)

	port := cfg.GetDefault(fmt.Sprintf("http.%s.listen_port", env), ":8080")
	readTimeOut := cfg.GetIntDefault("http.app_http_read_timeout", 60)
	writeTimeOut := cfg.GetIntDefault("http.app_http_write_timeout", 60)
	log2.Println("|-----------------------------------|")
	log2.Println("|            bingo-gin              |")
	log2.Println("|-----------------------------------|")
	log2.Println("|  Go Http Server Start Successful  |")
	log2.Println("|    Port" + port + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "        |")
	log2.Println("|-----------------------------------|")
	log2.Println("")
	server := &http.Server{
		Addr:           port,
		Handler:        ginApplication,
		ReadTimeout:    time.Duration(readTimeOut) * time.Second,
		WriteTimeout:   time.Duration(writeTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go handelStop(server)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger := log.LoggerHelp()
		logger.Errorf("HTTP server listen: %s\n", err)
	}
}



func (g *GinServerStart) Stop(ctx starters.StarterContext) {
	stopChan <- true
	return
}

func handelStop(server *http.Server) {
	<- stopChan
	log2.Println("Shutdown Http Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log2.Fatal("Http Server Shutdown:", err)
	}
	log2.Println("Http Server exiting")
	return
}


















