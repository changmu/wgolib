// Package wlog 封装常用的第三方日志库logrus初始化
package wlog

import (
	"fmt"
	"path/filepath"
	"runtime"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/changmu/wgolib/wlog/consts"
	"github.com/changmu/wgolib/wlog/opt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func init() {
	InitLogger()
}

// InitLogger 初始化logger
func InitLogger(optFuncs ...opt.LogOptFunc) {
	logOpt := opt.NewLogOpt(optFuncs...)
	Init(logOpt)
}

// Init 初始化logger
func Init(opt *opt.LogOpt) {
	logger := logrus.StandardLogger()
	// 配置日志等级
	logger.SetLevel(convertLogLevel(opt.LogLevel))
	// 配置滚动日志
	hook := NewLfsHook(opt.FileName, opt.KeepDays)
	logger.Hooks = logrus.LevelHooks{} // 修正重复初始化的问题
	logger.AddHook(hook)
	// 配置日志格式
	logger.SetFormatter(formatter())
	// 配置调用文件及行数
	logger.SetReportCaller(true)
}

func formatter() *nested.Formatter {
	fmtter := &nested.Formatter{
		HideKeys:        true,
		TimestampFormat: "2006-01-02 15:04:05.000", // 精确到毫秒
		CallerFirst:     true,
		NoColors:        true,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			funcInfo := runtime.FuncForPC(frame.PC)
			if funcInfo == nil {
				return "error during runtime.FuncForPC"
			}
			fullPath, line := funcInfo.FileLine(frame.PC)
			return fmt.Sprintf(" [%v:%v]", filepath.Base(fullPath), line)
		},
	}
	return fmtter
}

// NewLfsHook 设置滚动日志
func NewLfsHook(logName string, leastDay int) logrus.Hook {
	writer, err := rotatelogs.New(
		// 日志文件
		logName+".%Y%m%d.log",
		rotatelogs.WithRotationCount(uint(leastDay)), //只保留最近的N个日志文件
	)
	if err != nil {
		panic(err)
	}

	// 可设置按不同level创建不同的文件名
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.TraceLevel: writer,
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, formatter())

	return lfsHook
}

func convertLogLevel(level consts.ELogLevel) logrus.Level {
	switch level {
	case consts.LogLevelTrace:
		return logrus.TraceLevel
	case consts.LogLevelDebug:
		return logrus.DebugLevel
	case consts.LogLevelInfo:
		return logrus.InfoLevel
	case consts.LogLevelError:
		return logrus.ErrorLevel
	case consts.LogLevelFatal:
		return logrus.FatalLevel
	default:
		return logrus.DebugLevel
	}
}
