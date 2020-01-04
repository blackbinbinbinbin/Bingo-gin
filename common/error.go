package common

import (
	"fmt"
	"runtime"
)

type Error struct {
	Code  int
	Msg   string
	Where string
}

func (e *Error) Error() string {
	return fmt.Sprintf("code = %d ; msg = %s ; where = %s", e.Code, e.Msg, e.Where)
}

func NewError(code int, msg string) *Error {
	// 获取代码位置
	_, file, line, _ := runtime.Caller(1)
	where := fmt.Sprintf("%s:%d", file, line)

	err := &Error{Code:code, Msg:msg, Where: where}
	return err
}

func Warp(err error, msg string) *Error {
	var where string
	var code int
	switch t:=err.(type){
	case *Error:
		// 继承where和code
		where = t.Where
		code = t.Code
		// 拼接上之前的错误
		msg = msg + ":: " + t.Msg
	default:
		_, file, line, _ := runtime.Caller(1)
		code = CodeSystemError
		msg = msg + ":: " + err.Error()
		where = fmt.Sprintf("%s:$d", file, line)
	}

	error := &Error{Code:code, Msg:msg, Where: where}
	return error
}
