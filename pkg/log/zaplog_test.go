package log

import (
	"testing"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap"
	"fmt"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/log"
)

func TestLogger(t *testing.T) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		//StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
	}

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)

	config := zap.Config{
		Level:            atom,                                                // 日志级别
		Development:      true,                                                // 开发模式，堆栈跟踪
		Encoding:         "json",                                              // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                                       // 编码器配置
		OutputPaths:      []string{"stdout"},         // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}

	// 构建日志
	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("log 初始化失败: %v", err))
	}

	// 普通调用
	zplog := NewZapLogger(logger)
	log.Debug(zplog).Log(MSG, "test debug")
	log.Info(zplog).Log(MSG, "test info")
	log.Warn(zplog).Log(MSG, "test warn")
	log.Error(zplog).Log(MSG, "test error")


	// Helper调用
	logHleper := log.NewHelper("test env", zplog)
	logHleper.Debug("test debug")
	logHleper.Info("test info")
	logHleper.Warn("test warn")
	logHleper.Error("test error")
}