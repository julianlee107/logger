package logger

import "os"

type LogFile struct {
	*LogBase
	filename string
	file     *os.File
}

func (f *LogFile) Init() error {
	var err error
	f.file, err = os.OpenFile(f.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0677)
	return err
}

func NewLogFile(level int, filename, module string) Logger {
	logger := &LogFile{
		filename: filename,
	}
	logger.LogBase = &LogBase{level: level, module: module}
	return logger
}

func (f *LogFile) LogDebug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	logData := f.formatLogger(LogLevelDebug, f.module, format, args...)
	f.writeLog(f.file, logData)

}

func (f *LogFile) LogTrace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	logData := f.formatLogger(LogLevelTrace, f.module, format, args...)
	f.writeLog(f.file, logData)

}

func (f *LogFile) LogInfo(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	logData := f.formatLogger(LogLevelInfo, f.module, format, args...)
	f.writeLog(f.file, logData)

}

func (f *LogFile) LogError(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	logData := f.formatLogger(LogLevelError, f.module, format, args...)
	f.writeLog(f.file, logData)
}

func (f *LogFile) LogWarn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	logData := f.formatLogger(LogLevelWarn, f.module, format, args...)
	f.writeLog(f.file, logData)
}

func (f *LogFile) LogFatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	logData := f.formatLogger(LogLevelFatal, f.module, format, args...)
	f.writeLog(f.file, logData)

}

func (f *LogFile) SetLevel(level int) {
	f.level = level
}

func (f *LogFile) Close() {
	if f.file != nil {
		f.file.Close()
	}
}
