package consts

// ELogType 底层日志库类型
type ELogType int

const (
	LogTypeLogrus ELogType = 1
)

type ELogLevel int

const (
	LogLevelTrace = 0
	LogLevelDebug = 1
	LogLevelInfo  = 2
	LogLevelError = 3
	LogLevelFatal = 4
)
