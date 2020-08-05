package logger

type Logger interface {
	Init() error
	LogDebug(format string, args ...interface{})
	LogTrace(format string, args ...interface{})
	LogInfo(format string, args ...interface{})
	LogWarn(format string, args ...interface{})
	LogError(format string, args ...interface{})
	LogFatal(format string, args ...interface{})
	Close()
	SetLevel(level int)
}
