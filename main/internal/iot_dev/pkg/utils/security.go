package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// 使用aes gcm加密信息
func GcmEncrypt(plaintext []byte, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// 使用aes gcm解密信息
func GcmDecrypt(ciphertext []byte, key []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}

// 对字符串进行sha256加密，取前16个字符
func encryptKey(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	bs := fmt.Sprintf("%x", h.Sum(nil))
	return []byte(bs)[:16]
}

func Encrypt(src []byte, mac string) (crypted []byte, err error) {
	return GcmEncrypt(src, encryptKey(mac))
}

func Decrypt(crypted []byte, mac string) (origData []byte, err error) {
	return GcmDecrypt(crypted, encryptKey(mac))
}

// EncryptAuthPassword 账号密码加密
func EncryptAuthPassword(rawPassword string, saltKey string) (string, error) {
	bytes := []byte(rawPassword)
	encryptPassword, err := Encrypt(bytes, saltKey)
	if err != nil {
		return "", err
	}
	afterPassword := base64.StdEncoding.EncodeToString(encryptPassword)
	return afterPassword, nil
}

// DecryptAuthPassword 账号密码解密
func DecryptAuthPassword(encryptPassword string, saltKey string) (string, error) {
	bytesPassword, err := base64.StdEncoding.DecodeString(encryptPassword)
	if err != nil {
		return "", err
	}
	bytesPassword, err = Decrypt(bytesPassword, saltKey)
	if err != nil {
		return "", err
	}
	return string(bytesPassword), nil
}
