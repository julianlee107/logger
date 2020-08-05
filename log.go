package logger

// 定义全局默认日志
var log = NewLogger(LogTypeFile, LogLevelDebug, "", "default")

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

//封装接口
func Init(logType, level int, filename, module string) (err error) {

	log = NewLogger(logType, level, filename, module)
	return log.Init()
}

func LogDebug(format string, args ...interface{}) {
	log.LogDebug(format, args...)
}

func LogTrace(format string, args ...interface{}) {
	log.LogTrace(format, args...)
}

func LogInfo(format string, args ...interface{}) {
	log.LogInfo(format, args...)
}

func LogWarn(format string, args ...interface{}) {
	log.LogWarn(format, args...)
}

func LogError(format string, args ...interface{}) {
	log.LogError(format, args...)
}

func LogFatal(format string, args ...interface{}) {
	log.LogFatal(format, args...)
}

func SetLevel(level int) {
	log.SetLevel(level)
}

func Close() {
	log.Close()
}
