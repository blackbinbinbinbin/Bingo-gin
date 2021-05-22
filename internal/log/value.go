package log

import (
	"runtime"
	"strings"
	"strconv"
	"time"
)


var (
	// DefaultCaller is a Valuer that returns the file and line.
	DefaultCaller = Caller(3)

	// DefaultTimestamp is a Valuer that returns the current wallclock time.
	DefaultTimestamp = Timestamp(time.RFC3339)
)


// Valuer is returns a log value function.
type Valuer func() interface{}

// Value return the function value.
func Value(v interface{}) interface{} {
	if v, ok := v.(Valuer); ok {
		return v()
	}
	return v
}


// valuer：print caller stack
func Caller(depth int) Valuer {
	return func() interface{} {
		_, file, line, _ := runtime.Caller(depth)

		idx := strings.LastIndexByte(file, '/')
		return file[idx+1:] + ":" + strconv.Itoa(line)
	}
}


// Timestamp returns a timestamp Valuer with a custom time format.
func Timestamp(layout string) Valuer {
	return func() interface{} {
		return time.Now().Format(layout)
	}
}


// if value is valuer func. bind the func in kvs
func bindValues(keyvals []interface{}) {
	for i := 1; i < len(keyvals); i += 2 {
		if v, ok := keyvals[i].(Valuer); ok {
			keyvals[i] = v()
		}
	}
}


// judge the value is function or not
func containsValuer(keyvals []interface{}) bool {
	for i := 1; i < len(keyvals); i += 2 {
		if _, ok := keyvals[i].(Valuer); ok {
			return true
		}
	}
	return false
}