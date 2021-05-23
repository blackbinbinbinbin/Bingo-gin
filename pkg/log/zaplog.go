package log

import (
	"go.uber.org/zap"
	"github.com/blackbinbinbinbin/Bingo-gin/internal/log"
)

const MSG  = "msg"

type zapLogger struct {
	log  		*zap.Logger
	fields 		[]zap.Field
}


func (l *zapLogger) Log(kv ...interface{}) error {
	if len(kv) == 0 {
		return nil
	}
	if len(kv)%2 != 0 {
		kv = append(kv, "")
	}

	zpfields := []zap.Field{}
	level := log.LevelInfo
	msg := ""
	for i := 0; i < len(kv); i += 2 {
		if kv[i] == log.LevelKey {
			level = kv[i+1].(log.Level)
			continue
		}
		if kv[i].(string) == MSG {
			msg = kv[i+1].(string)
			continue
		}

		zpfields = append(zpfields, zap.Any(kv[i].(string), kv[i+1]))
	}
	l.log.With(zpfields...)

	switch level {
	case log.LevelDebug:
		l.log.Debug(msg, zpfields...)
	case log.LevelError:
		l.log.Error(msg, zpfields...)
	case log.LevelInfo:
		l.log.Info(msg, zpfields...)
	case log.LevelWarn:
		l.log.Warn(msg, zpfields...)
	default :
		l.log.Info(msg, zpfields...)
	}
	return nil
}


func NewZapLogger(logger *zap.Logger) log.Logger {
	return &zapLogger{
		log: logger,
	}
}