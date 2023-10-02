package utils

import (
	"math/rand"
	"time"
)

// 生成指定位数的随机数
func GenerateRandomID(length int) string {
	const charset = "0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	id := make([]byte, length)
	for i := range id {
		id[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(id)
}
