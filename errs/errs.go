package errs

import "fmt"

// WError 自定义错误
type WError struct {
	Code int
	Msg  string
}

// Error 错误标准接口
func (e WError) Error() string {
	if e.Code == RetSuccess {
		return ""
	}
	return fmt.Sprintf("code:%v,msg:%v", e.Code, e.Msg)
}

// New 创建error
func New(code int, format string, args ...interface{}) error {
	return WError{
		Code: code,
		Msg:  fmt.Sprintf(format, args...),
	}
}
