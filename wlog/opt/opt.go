package opt

import "github.com/changmu/wgolib/wlog/consts"

// LogOpt 日志配置项
type LogOpt struct {
	LogType  consts.ELogType
	FileName string // 比如wlog，不需要带.log后缀
	LogLevel consts.ELogLevel
	KeepDays int // 日志保留天数
}

type LogOptFunc func(opt *LogOpt)

// NewLogOpt 创建配置
func NewLogOpt(optFuncs ...LogOptFunc) *LogOpt {
	logOpt := &LogOpt{
		LogType:  consts.LogTypeLogrus,
		FileName: "wlog",
		LogLevel: consts.LogLevelDebug,
		KeepDays: 7,
	}
	for _, optFunc := range optFuncs {
		optFunc(logOpt)
	}
	return logOpt
}

// WithLogType 配置logType
func WithLogType(logType consts.ELogType) LogOptFunc {
	return func(opt *LogOpt) {
		opt.LogType = logType
	}
}

// WithLogLevel 配置日志等级
func WithLogLevel(logLevel consts.ELogLevel) LogOptFunc {
	return func(opt *LogOpt) {
		opt.LogLevel = logLevel
	}
}

// WithFileName 配置日志文件名
func WithFileName(fileName string) LogOptFunc {
	return func(opt *LogOpt) {
		opt.FileName = fileName
	}
}

// WithKeepDays 配置日志保留天数
func WithKeepDays(keepDays int) LogOptFunc {
	return func(opt *LogOpt) {
		opt.KeepDays = keepDays
	}
}
