package logger

import (
	"fmt"
	"os"
	"time"
)

type Log struct {
	timeStr  string
	levelStr string
	module   string
	filename string
	funcName string
	lineNo   int
	data     string
}

type LogBase struct {
	level  int
	module string
}

func (l *LogBase) writeLog(file *os.File, log *Log) {
	fmt.Fprintf(file, "%s [%s]%s (%s:%s:%d) %s \n",
		log.timeStr, log.levelStr, log.module, log.filename, log.funcName, log.lineNo, log.data)
}

func (l *LogBase) formatLogger(level int, module, format string, args ...interface{}) *Log {
	data := fmt.Sprintf(format, args...)
	fmt.Printf(format, args...)
	filename, funcName, lineNo := GetLineInfo(3)
	return &Log{
		timeStr:  time.Now().Format("2006-01-02 15:04:05.000"),
		levelStr: getLevelStr(level),
		module:   module,
		filename: filename,
		funcName: funcName,
		lineNo:   lineNo,
		data:     data,
	}
}
