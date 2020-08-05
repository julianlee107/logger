package main

import (
	"flag"
	"fmt"
	"github.com/julianlee107/logger"
)

func logic(logger logger.Logger) {
	for {
		logger.LogDebug("jiuquhe %s %s", "123", "chongchongchong")
		logger.LogTrace("dads2")
		logger.LogInfo("dads3")
		logger.LogWarn("dads4")
		logger.LogError("sss")
		logger.LogFatal("sss")
	}

}

func testGetLine() {
	filename, funcName, lineNo := logger.GetLineInfo(2)
	fmt.Printf("filename=%s,funcname=%s,linenum=%d\n", filename, funcName, lineNo)
}

func main() {
	testGetLine()
	var logTypeStr string
	flag.StringVar(&logTypeStr, "type", "file", "please input a logger type")
	flag.Parse()

	var logType int
	if logTypeStr == "file" {
		logType = logger.LogTypeFile
	} else {
		logType = logger.LogTypeConsole
	}
	log := logger.NewLogger(logType, logger.LogLevelDebug, "./logger.log", "logger_example")
	err := log.Init()
	if err != nil {
		fmt.Printf("init error:%v\n", err)
		return
	}
	logic(log)
	log.Close()
}
