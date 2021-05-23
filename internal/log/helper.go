package log

import "fmt"

type Helper struct {
	debug 	Logger
	info	Logger
	warn 	Logger
	err 	Logger
}

// new a logger helper. config a third party pkg.
func NewHelper(name string, logger Logger) *Helper {
	logger = With(logger, "module", name)

	return &Helper{
		debug: Debug(logger),
		info:  Info(logger),
		warn:  Warn(logger),
		err:   Error(logger),
	}
}

// Debug logs a message at debug level.
func (h *Helper) Debug(a ...interface{}) {
	h.debug.Log("msg", fmt.Sprint(a...))
}

// Debugf logs a message at debug level.
func (h *Helper) Debugf(format string, a ...interface{}) {
	h.debug.Log("msg", fmt.Sprintf(format, a...))
}

// Debugw logs a message at debug level.
func (h *Helper) Debugw(kv ...interface{}) {
	h.debug.Log(kv...)
}

// Info logs a message at info level.
func (h *Helper) Info(a ...interface{}) {
	h.info.Log("msg", fmt.Sprint(a...))
}

// Infof logs a message at info level.
func (h *Helper) Infof(format string, a ...interface{}) {
	h.info.Log("msg", fmt.Sprintf(format, a...))
}

// Infow logs a message at info level.
func (h *Helper) Infow(kv ...interface{}) {
	h.info.Log(kv...)
}

// Warn logs a message at warn level.
func (h *Helper) Warn(a ...interface{}) {
	h.warn.Log("msg", fmt.Sprint(a...))
}

// Warnf logs a message at warnf level.
func (h *Helper) Warnf(format string, a ...interface{}) {
	h.warn.Log("msg", fmt.Sprintf(format, a...))
}

// Warnw logs a message at warnf level.
func (h *Helper) Warnw(kv ...interface{}) {
	h.warn.Log(kv...)
}

// Error logs a message at error level.
func (h *Helper) Error(a ...interface{}) {
	h.err.Log("msg", fmt.Sprint(a...))
}

// Errorf logs a message at error level.
func (h *Helper) Errorf(format string, a ...interface{}) {
	h.err.Log("msg", fmt.Sprintf(format, a...))
}

// Errorw logs a message at error level.
func (h *Helper) Errorw(kv ...interface{}) {
	h.err.Log(kv...)
}



