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

func NewLogger(logType, level int, filename, module string) Logger {
	var logger Logger
	switch logType {
	case LogTypeFile:
		logger = NewLogFile(level, filename, module)
	case LogTypeConsole:
		logger = NewLogConsole(level, module)
	default:
		logger = NewLogFile(level, filename, module)
	}
	return logger
}
