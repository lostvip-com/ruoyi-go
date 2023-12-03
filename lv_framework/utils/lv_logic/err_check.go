package lv_logic

import (
	"encoding/json"
	"log"
	"lostvip.com/web/dto"
	"runtime"
)

// HasError 错误断言
// 当 error 不为 nil 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
// 若 msg 为空，则默认为 error 中的内容
func Assert1(conditionYes bool, msg string) {
	if conditionYes {
		var res dto.Resp
		res.Msg = msg
		res.Code = 1
		json, _ := ToJsonStr(res)
		panic(json)
	}
}
func ToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

// HasError 错误断言
// 当 error 不为 nil 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
// 若 msg 为空，则默认为 error 中的内容
func HasErrorMsg(err error, msg string) {
	if err != nil {
		if msg == "" {
			msg = err.Error()
		}
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%v error: %#v", file, line, err)
		var res dto.Resp
		res.Msg = msg
		res.Code = 1
		panic(res)
	}
}

func HasError1(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%v error: %#v", file, line, err)
		var res dto.Resp
		res.Msg = err.Error()
		res.Code = 1
		panic(res)
	}
}
