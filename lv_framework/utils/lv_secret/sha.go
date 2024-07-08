// ////////////////////////////////////////////////////////////////
//
// 以下Golang代码的加密结果与Java语言结果一致，需要注意结果大小写问题。
// ////////////////////////////////////////////////////////////////
package lv_secret

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// hmacsha256验证
func HMAC_SHA256(src, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))

}

// hmacsha512验证
func HMAC_SHA512(src, key string) string {
	m := hmac.New(sha512.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

func HMAC_SHA1(src, key string) string {
	m := hmac.New(sha1.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

// sha256验证

func SHA256Str(src string) string {

	h := sha256.New()

	h.Write([]byte(src)) // 需要加密的字符串为

	// fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果

	return hex.EncodeToString(h.Sum(nil))

}

// sha512验证

func SHA512Str(src string) string {

	h := sha512.New()

	h.Write([]byte(src)) // 需要加密的字符串为

	// fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果

	return hex.EncodeToString(h.Sum(nil))

}
