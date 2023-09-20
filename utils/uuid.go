package utils

import (
	"github.com/google/uuid"
	"math/rand"
	"strings"
	"time"
)

// NewUUID new uuid
func NewUUID() string {
	return uuid.New().String()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewBizID 生成业务Id
// bizCode  业务编码
// suffixLength 结尾随机字符长度
func NewBizID(bizCode string, suffixLength int) string {
	time := Timestamp()
	suffix := randStr(suffixLength)
	return bizCode + time + suffix
}

// Timestamp 时间戳，年份取2位，直到毫秒
func Timestamp() string {
	timeStr := time.Now().Format("20060102150405.0000")
	return strings.ReplaceAll(timeStr, ".", "")[2:]
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 随机字符串
func randStr(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
