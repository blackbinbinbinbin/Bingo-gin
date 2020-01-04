package main

import (
	"Bingo-gin/router"
	"Bingo-gin/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"strconv"
	"fmt"
	"os"
	"log"
	"os/signal"
	"context"
	"io"
	"path/filepath"
	"strings"
)

func main() {
	/**
	获取项目配置
	 */
	cfg := common.NewConfig()
	appMode := cfg.GetValue("", "APP_ENV_MODE")
	gin.SetMode(appMode)

	/**
	设置日志格式，和输出重定向到文件，并且清理过期日志
	 */
	log.SetFlags(log.LstdFlags | log.Llongfile |log.LUTC)
	consoleLogFileName := cfg.GetValue("console.log", "APP_CONSOLE_LOG_FILE")
	now := time.Now()
	logFile := fmt.Sprintf("%s/%s.%s.log", common.ProjectLogPath, consoleLogFileName, now.Format("20060102"))
	logWriter, e := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
	if e != nil {
		log.Println("打开日志文件错误")
	} else {
		gin.DefaultWriter = io.Writer(logWriter)
		log.SetOutput(gin.DefaultWriter)
	}
	checkLastDate := time.Now().AddDate(0, 0, -7)
	checkLastDateInt, _ := strconv.Atoi(checkLastDate.Format("20060102"))
	filepath.Walk(common.ProjectLogPath, func (path string, info os.FileInfo, err error) error {
		if pos := strings.Index(path, consoleLogFileName); pos != -1 {
			logFileSliceString := strings.Split(path, ".")
			for _, str := range logFileSliceString {
				date , _ := strconv.Atoi(str)
				if date != 0 && date < checkLastDateInt {
					os.Remove(path)
				}
			}
		}
		return nil
	})

	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		fmt.Println("in go")
		c.String(200, "Hello Bingo-gin")
	})

	/**
	初始化路由
	 */
	router.InitRouter(engine)

	/**
	定义http服务
	 */
	port := cfg.GetSectionValue(appMode, "LISTEN_PORT", ":8080")
	readTimeOutConfig := cfg.GetValue("60", "APP_HTTP_READ_TIMEOUT")
	readTimeOut, _ := strconv.ParseInt(readTimeOutConfig, 10, 64)
	writeTimeOutConfig := cfg.GetValue("60", "APP_HTTP_WRITE_TIMEOUT")
	writeTimeOut, _ := strconv.ParseInt(writeTimeOutConfig, 10, 64)
	log.Println("|-----------------------------------|")
	log.Println("|            bingo-gin              |")
	log.Println("|-----------------------------------|")
	log.Println("|  Go Http Server Start Successful  |")
	log.Println("|    Port" + port + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "        |")
	log.Println("|-----------------------------------|")
	log.Println("")
	server := &http.Server{
		Addr:           port,
		Handler:        engine,
		ReadTimeout:    time.Duration(readTimeOut) * time.Second,
		WriteTimeout:   time.Duration(writeTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()

	// 等待中断信号以关闭服务器-设置 5 秒的超时时间
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}