package lv_err

import (
	"fmt"
	"runtime/debug"
)

func PrintStackTrace(err error) string {
	s := string(debug.Stack()) //打印错误堆栈
	msg := fmt.Sprintf("err=%v, stack=%s\n", err, s)
	fmt.Println(msg)
	return msg
}
