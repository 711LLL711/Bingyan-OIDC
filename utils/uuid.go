package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
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

func GenerateVerificationCode() string {
	// 生成一个随机的 UUID
	uuidObj := uuid.New()

	// 将 UUID 转换为字符串，并移除其中的短划线 "-"
	uuidStr := strings.ReplaceAll(uuidObj.String(), "-", "")

	// 提取其中的字符和数字部分
	verificationCode := ""
	for _, char := range uuidStr {
		if (char >= '0' && char <= '9') || (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			verificationCode += string(char)
		}
	}

	// 返回验证码
	return verificationCode
}
