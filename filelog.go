package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type LogFile struct {
	*LogBase
	filename string
	file     *os.File

	//	异步写入
	logChan chan *Log
	wg      *sync.WaitGroup
	// 日志切分，按天
	curDay int
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
	logger.curDay = time.Now().Day()
	logger.wg = &sync.WaitGroup{}
	// 管道初始化
	logger.logChan = make(chan *Log, 10000)

	//异步写日志到磁盘
	logger.wg.Add(1)
	go logger.syncLog()
	return logger
}

func (f *LogFile) syncLog() {
	for data := range f.logChan {
		f.writeLog(f.file, data)
	}
	f.wg.Done()
}

//日志切分
func (f *LogFile) splitLog() {
	now := time.Now()
	if now.Day() == f.curDay {
		return
	}
	f.curDay = now.Day()
	f.file.Sync()
	f.file.Close()

	newFile := fmt.Sprintf("%s-%04d-%02d-%02d", f.filename, now.Year(), now.Month(), now.Day())
	os.Rename(f.filename, newFile)
	f.Init()
}

// 日志写入到chan中
func (f *LogFile) writeToChan(level int, module, format string, args ...interface{}) {
	logData := f.formatLogger(level, f.module, format, args...)
	f.logChan <- logData
}

func (f *LogFile) LogDebug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	f.writeToChan(LogLevelDebug, f.module, format, args...)

}

func (f *LogFile) LogTrace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	f.writeToChan(LogLevelTrace, f.module, format, args...)

}

func (f *LogFile) LogInfo(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	f.writeToChan(LogLevelInfo, f.module, format, args...)

}

func (f *LogFile) LogError(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	f.writeToChan(LogLevelError, f.module, format, args...)
}

func (f *LogFile) LogWarn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	f.writeToChan(LogLevelWarn, f.module, format, args...)
}

func (f *LogFile) LogFatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	f.writeToChan(LogLevelFatal, f.module, format, args...)

}

func (f *LogFile) SetLevel(level int) {
	f.level = level
}

func (f *LogFile) Close() {
	if f.logChan != nil {
		close(f.logChan)
	}
	f.wg.Wait()
	if f.file != nil {
		f.file.Sync()
		f.file.Close()
	}
}
