// ////////////////////////////////////////////////////////////////
//
// 以下Golang代码的加密结果与Java语言结果一致，需要注意结果大小写问题。
// ////////////////////////////////////////////////////////////////
package lv_secret

import (
	"bytes"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
