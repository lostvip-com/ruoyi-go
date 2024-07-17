package utils

import (
	"fmt"
	"hash/crc32"
	rand2 "math/rand"
	"time"

	"github.com/google/uuid"
)

func GenUUID() string {
	uuidHash := crc32.ChecksumIEEE([]byte(uuid.New().String()))
	return fmt.Sprint(uuidHash)
}

func RandomNum() string {
	return fmt.Sprintf("%08v", rand2.New(rand2.NewSource(time.Now().UnixNano())).Int31n(100000000))
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func GenerateDeviceSecret(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand2.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand2.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func GenerateSaltKey() string {
	return uuid.New().String()
}
