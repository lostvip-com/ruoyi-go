package lv_reflect

import (
	"runtime"
	"strings"
)

func GetMethodName() string {
	frame, _, _, _ := runtime.Caller(1)
	// 获取方法名
	funcName := runtime.FuncForPC(frame).Name()
	// 获取包名
	funcName = funcName[strings.LastIndex(funcName, ".")+1:]
	return funcName
}
