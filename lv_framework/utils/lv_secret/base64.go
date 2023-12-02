// ////////////////////////////////////////////////////////////////
//
// 以下Golang代码的加密结果与Java语言结果一致，需要注意结果大小写问题。
// ////////////////////////////////////////////////////////////////
package lv_secret

import (
	"encoding/base64"
)

// base编码

func BASE64EncodeStr(src string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(src)))
}

// base解码

func BASE64DecodeStr(src string) string {
	a, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	return string(a)
}
