// ////////////////////////////////////////////////////////////////
//
// 以下Golang代码的加密结果与Java语言结果一致，需要注意结果大小写问题。
// ////////////////////////////////////////////////////////////////
package lv_secret

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"strings"
)

var ivspec = []byte("0000000000000000")

func AESEncodeStr(src, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("key error1", err)
	}

	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, ivspec)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return hex.EncodeToString(crypted)

}

func AESDecodeStr(crypt, key string) string {
	crypted, err := hex.DecodeString(strings.ToLower(crypt))
	if err != nil || len(crypted) == 0 {
		fmt.Println("plain content empty")
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("key error1", err)
	}

	ecb := cipher.NewCBCDecrypter(block, ivspec)
	decrypted := make([]byte, len(crypted))
	ecb.CryptBlocks(decrypted, crypted)
	return string(PKCS5Trimming(decrypted))
}
