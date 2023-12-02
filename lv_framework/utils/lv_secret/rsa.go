// ////////////////////////////////////////////////////////////////
//
// 以下Golang代码的加密结果与Java语言结果一致，需要注意结果大小写问题。
// ////////////////////////////////////////////////////////////////
package lv_secret

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
)

func RsaEncrypt(src, key string) string {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return ""
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
		return ""
	}
	pub := pubInterface.(*rsa.PublicKey)
	crypted, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(src))
	if err != nil {
		panic(err)
		return ""
	}
	return hex.EncodeToString(crypted)

}
