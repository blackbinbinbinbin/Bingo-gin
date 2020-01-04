package logger

import (
	"github.com/gin-gonic/gin"
	"Bingo-gin/common"
	"github.com/sirupsen/logrus"
	"fmt"
	"time"
	"Bingo-gin/util/response"
	"bytes"
	"encoding/json"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}


func SetUp() gin.HandlerFunc {
	cfg := common.NewConfig()
	// 日志路径
	logFilePath := common.ProjectLogPath
	logFileName := cfg.GetValue("access.log", "APP_ACCESS_LOG_FILE")

	// 日志文件
	logFile := fmt.Sprintf("%s/%s.log", logFilePath, logFileName)

	// 实例化
	logger := logrus.New()

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})

	// 设置 rotatelogs
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		logFile + ".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(logFile),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})

	// 新增 Hook
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		responseBody := bodyLogWriter.body.String()

		var responseCode int
		var responseMsg string
		var responseData interface{}

		if responseBody != "" {
			res := response.Response{}
			err := json.Unmarshal([]byte(responseBody), &res)
			if err == nil {
				responseCode = res.Code
				responseMsg = res.Msg
				responseData = res.Data
			} else {
				responseCode = 0
				responseMsg = ""
				responseData = make([]string, 1)
			}
		}

		// 执行时间
		latencyTime := endTime.Sub(startTime).Microseconds()

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 请求数据
		reqPostData := c.Request.PostForm.Encode()

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.WithFields(logrus.Fields{
			"http_status_code"  : statusCode,
			"latency_time" : latencyTime,
			"client_ip"    : clientIP,
			"req_method"   : reqMethod,
			"req_uri"      : reqUri,
			"req_post_data": reqPostData,
			"response_time": endTime,
			"response_code": responseCode,
			"response_msg" : responseMsg,
			"response_data": responseData,
			"response_body": responseBody,
		}).Info()
	}
}