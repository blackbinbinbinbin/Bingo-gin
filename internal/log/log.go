package log

import (
	"log"
)

var (
	// DefaultLogger is default logger.
	DefaultLogger Logger = NewStdLogger(log.Writer())
)

// Logger is a logger interface.
type Logger interface {
	Log(kv ...interface{}) error
}

// contest is Logger
// context deals with assembling key and value
// prefix is the prefix kv
// logs is a array of log. It call log's Log func.
// hasValuer means kvs's value is a function
type context struct {
	logs 		[]Logger
	prefix 		[]interface{}
	hasValuer	bool
}

// this context achieve the Logger interface
func (c *context) Log(kv ...interface{}) error {
	kvs := make([]interface{}, 0, len(c.prefix) + len(kv))
	kvs = append(kvs, c.prefix...)
	if c.hasValuer {
		// bind Value
		bindValues(kvs)
	}

	for _,l := range c.logs {
		if err := l.Log(kvs...); err != nil {
			return err
		}
	}

	return nil
}


func With(l Logger, kv ...interface{}) Logger {
	if c, ok := l.(*context); ok {
		kvs := make([]interface{}, 0, len(c.prefix)+len(kv))
		kvs = append(kvs, kv...)
		kvs = append(kvs, c.prefix...)
		return &context{
			logs:      c.logs,
			prefix:    kvs,
			hasValuer: containsValuer(kvs),
		}
	}
	return &context{logs: []Logger{l}, prefix: kv, hasValuer: containsValuer(kv)}
}


func MultiLogger(logs ...Logger) Logger {
	return &context{logs: logs}
}


// Debug returns a debug logger.
func Debug(log Logger) Logger {
	return With(log, LevelKey, LevelDebug)
}

// Info returns a info logger.
func Info(log Logger) Logger {
	return With(log, LevelKey, LevelInfo)
}

// Warn return a warn logger.
func Warn(log Logger) Logger {
	return With(log, LevelKey, LevelWarn)
}

// Error returns a error logger.
func Error(log Logger) Logger {
	return With(log, LevelKey, LevelError)
}