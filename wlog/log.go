// Package wlog 封装常用的第三方日志库
package wlog

import (
	"github.com/changmu/wgolib/wlog/consts"
	"github.com/changmu/wgolib/wlog/logrus"
	"github.com/changmu/wgolib/wlog/opt"
)

func init() {
	logger = NewLogger()
}

var logger Logger

// Logger 日志通用接口
type Logger interface {
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

// NewLogger 创建logger
func NewLogger(optFuncs ...opt.LogOptFunc) Logger {
	// 默认配置
	logOpt := opt.NewLogOpt(optFuncs...)
	// 创建日志
	return newLogger(logOpt)
}

func newLogger(opt *opt.LogOpt) Logger {
	switch opt.LogType {
	case consts.LogTypeLogrus:
		return logrus.New(opt)
	default:
		return nil
	}
}
