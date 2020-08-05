package logger

import "runtime"

func GetLineInfo(skip int) (filename, funcName string, lineNo int) {
	pc, filename, lineNo, ok := runtime.Caller(skip)
	if ok {
		fun := runtime.FuncForPC(pc)
		funcName = fun.Name()
	}
	return

}
