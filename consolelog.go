package logger

import "os"

type LogConsole struct {
	*LogBase
}

func (c *LogConsole) Init() error {
	return nil
}

func NewLogConsole(level int, module string) Logger {
	logger := &LogFile{}
	logger.LogBase = &LogBase{level: level, module: module}
	return logger
}

func (c *LogConsole) LogDebug(format string, args ...interface{}) {
	if c.level > LogLevelDebug {
		return
	}
	logData := c.formatLogger(LogLevelDebug, c.module, format)
	c.writeLog(os.Stdout, logData)

}

func (c *LogConsole) LogTrace(format string, args ...interface{}) {
	if c.level > LogLevelTrace {
		return
	}
	logData := c.formatLogger(LogLevelTrace, c.module, format)
	c.writeLog(os.Stdout, logData)

}

func (c *LogConsole) LogInfo(format string, args ...interface{}) {
	if c.level > LogLevelInfo {
		return
	}
	logData := c.formatLogger(LogLevelInfo, c.module, format)
	c.writeLog(os.Stdout, logData)

}

func (c *LogConsole) LogError(format string, args ...interface{}) {
	if c.level > LogLevelError {
		return
	}
	logData := c.formatLogger(LogLevelError, c.module, format)
	c.writeLog(os.Stdout, logData)
}

func (c *LogConsole) LogWarn(format string, args ...interface{}) {
	if c.level > LogLevelWarn {
		return
	}
	logData := c.formatLogger(LogLevelWarn, c.module, format)
	c.writeLog(os.Stdout, logData)
}

func (c *LogConsole) LogFatal(format string, args ...interface{}) {
	if c.level > LogLevelFatal {
		return
	}
	logData := c.formatLogger(LogLevelFatal, c.module, format)
	c.writeLog(os.Stdout, logData)

}

func (c *LogConsole) SetLevel(level int) {
	c.level = level
}

func (c *LogConsole) Close() {
}
