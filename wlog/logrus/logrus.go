package logrus

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

// Logger 封装logrus日志库
type Logger struct {
	imp *logrus.Logger
	opt *opt.LogOpt
}

// New 创建logger
// 由于CustomCallerFormatter没法返回包装后的函数调用者，因此这里返回裸的logrus.Logger
func New(opt *opt.LogOpt) *logrus.Logger {
	logger := logrus.New()
	// 配置日志等级
	logger.SetLevel(convertLogLevel(opt.LogLevel))
	// 配置滚动日志
	hook := NewLfsHook(opt.FileName, opt.KeepDays)
	logger.AddHook(hook)
	// 配置日志格式
	logger.SetFormatter(formatter())
	// 配置调用文件及行数
	logger.SetReportCaller(true)
	return logger
}

func formatter() *nested.Formatter {
	fmtter := &nested.Formatter{
		HideKeys:        true,
		TimestampFormat: "2006-01-02 15:04:05.000",
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

func (logger *Logger) Tracef(format string, args ...interface{}) {
	logger.imp.Tracef(format, args...)
}
func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.imp.Debugf(format, args...)
}
func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.imp.Infof(format, args...)
}
func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.imp.Errorf(format, args...)
}
func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.imp.Fatalf(format, args...)
}

func NewLfsHook(logName string, leastDay int) logrus.Hook {
	writer, err := rotatelogs.New(
		// 日志文件
		logName+".%Y%m%d.wlog",
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
