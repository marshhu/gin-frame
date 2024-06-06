package utils

import (
	"errors"
	"strconv"
	"strings"
)

// ToInt to int64
func ToInt(val string) (int64, error) {
	return strconv.ParseInt(val, 10, 64)
}

// ToInt32 to int32
func ToInt32(val string) (int, error) {
	t, err := strconv.ParseInt(val, 10, 32)
	return int(t), err
}

// ToFloat to float
func ToFloat(val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

// Int64ToString int64 to string
func Int64ToString(val int64) string {
	return strconv.FormatInt(val, 10)
}

// Float64ToString float to string
func Float64ToString(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}

// Split fen ci
func Split(val string, ch rune) []string {
	values := strings.FieldsFunc(val, func(r rune) bool {
		return r == ch
	})

	return values
}

// SplitToFloat split to float
func SplitToFloat(val string, ch rune) []float64 {
	var ret []float64
	values := Split(val, ch)
	for _, val := range values {
		if v, err := strconv.ParseFloat(val, 64); err == nil {
			ret = append(ret, v)
		}
	}
	return ret
}

// SplitToInt split to int
func SplitToInt(val string, ch rune) []int64 {
	var ret []int64
	values := Split(val, ch)
	for _, val := range values {
		if v, err := strconv.ParseInt(val, 10, 64); err == nil {
			ret = append(ret, v)
		}
	}
	return ret
}

// Trim trm
func Trim(val string) string {
	return strings.Trim(val, " ")
}

// ToBool to bool
func ToBool(val string) (bool, error) {
	return strconv.ParseBool(val)
}

// BoolToString bool to string
func BoolToString(val bool) string {
	return strconv.FormatBool(val)
}

// IsEmpty 为空判断
func IsEmpty(str string) (isEmpty bool) {
	return len(Trim(str)) == 0
}

// ArrayInt64ToString []int64 to string
func ArrayInt64ToString(elems []int64, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return Int64ToString(elems[0])
	}

	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(Int64ToString(elems[i]))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(Int64ToString(elems[0]))
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(Int64ToString(s))
	}
	return b.String()
}

// ReplaceStrEmptyAndLine ReplaceEmpty
func ReplaceStrEmptyAndLine(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}

// SubString SubString截取指定长度的字符串
func SubString(value string, length int) (string, error) {
	var maxIndex = len(value) - 1
	runes := []rune(value)
	if length-1 > maxIndex {
		return "", errors.New("length is out of max length of the string")
	}
	return string(runes[0 : length-1]), nil
}

// MarkUpSuffix 按照特定字符自动尾部补全
func MarkUpSuffix(value string, charSet string, count int) string {
	return value + strings.Repeat(charSet, count)
}

// StringsSplitLimit 按批次分隔字符串，每批次字符串的值用sep分隔符分隔，每批次字符串的值最多数量是limitInt
// 如：arr = []string{"1","2","3","4","5","6","7"}
// sep = ";"
// limitInt = 2
// 返回结果：[]string{"1;2", "3;4", "5;6", "7"}
func StringsSplitLimit(arr []string, sep string, limitInt int) []string {
	var ret []string

	if arr == nil || len(arr) == 0 {
		return nil
	}

	// 每批次字符串的值最多数量 == 0 时，或者数组长度小于每批次字符串的值最多数量时，返回源字符串数组转成的字符串
	if limitInt == 0 || len(arr) <= limitInt {
		ret = append(ret, strings.Join(arr, sep))
		return ret
	}

	tmpStr := ""
	for i, o := range arr {
		if i == 0 {
			tmpStr = o
		} else if i%limitInt != 0 {
			tmpStr += sep + o
		} else {
			ret = append(ret, tmpStr)
			tmpStr = o
		}
	}
	// 将最后一批组装
	if len(tmpStr) > 0 {
		ret = append(ret, tmpStr)
	}
	return ret
}

// Contains 字符串切片包含字符串，检查数据包含关系
func Contains(array []string, item string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == item {
			return true
		}
	}
	return false
}

func AppendZero(str string, strLen int) string {
	if len(str) >= strLen {
		return str
	}
	charArray := []rune(str)
	for i := 0; i < strLen-len(str); i++ {
		charArray = append(charArray, '0')
	}
	return string(charArray)
}
