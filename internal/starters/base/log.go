package base

import (
	"github.com/blackbinbinbinbin/Bingo-gin/internal/log"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/starters"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap"
	"fmt"
	log2 "github.com/blackbinbinbinbin/Bingo-gin/pkg/log"
	"os"
)

const (
	LOG_OUT_TYPE_STD = "stdout"
)

var logger *log.Helper

func LoggerHelp() *log.Helper {
	starters.Check(logger)
	return logger
}



type LogStarter struct {
	starters.BaseStarter
}
func (l *LogStarter) Init(ctx starters.StarterContext) {
	// 获取配置
	props := ctx.Props()
	outfile := props.GetDefault("log.output_path", LOG_OUT_TYPE_STD)
	if outfile != LOG_OUT_TYPE_STD {
		_, err := os.Lstat(outfile)
		if os.IsNotExist(err) {
			outfile = LOG_OUT_TYPE_STD
		}
	}

	// 初始化 zap logger
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
	if props.GetBoolDefault("log.open_stacktrace", false) {
		encoderConfig.StacktraceKey = "stacktrace"
	}

	// 设置日志级别
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)

	config := zap.Config{
		Level:            atom,                                                // 日志级别
		Development:      true,                                                // 开发模式，堆栈跟踪
		Encoding:         "json",                                              // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                                       // 编码器配置
		OutputPaths:      []string{outfile},         // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{outfile},
	}

	// 构建日志
	zplogger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("log 初始化失败: %v", err))
	}

	module := props.GetDefault("module", "Bingo-gin")
	logger = log.NewHelper(module, log2.NewZapLogger(zplogger))
}