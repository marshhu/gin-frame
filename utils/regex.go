package utils

import (
	"regexp"
)

// StringRegex 正则
func StringRegex(str string, strType string) bool {
	var pattern = ""
	switch strType {
	case "mobile":
		pattern = "^(1[3,4,5,6,7,8,9][0-9])\\d{8}$"
	case "email":
		pattern = `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	case "datetime":
		pattern = `^([0-9]{4}[-/][0-9]{2}[-/][0-9]{2})( [0-9]{2}:[0-9]{2}(:[0-9]{2})?)?$`
	default:
		return false
	}
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(str)
}

func IsNumeric(s string) bool {
	// 正则表达式匹配整数或小数
	reIntOrDecimal := regexp.MustCompile(`^(-?\d+)(\.\d+)?$`)
	return reIntOrDecimal.MatchString(s)
}

// IsFraction 判断字符串是否匹配分数形式
func IsFraction(s string) bool {
	// 正则表达式匹配简单的分数形式
	reFraction := regexp.MustCompile(`^\d+/\d+$`)
	// 检查字符串是否匹配分数形式
	return reFraction.MatchString(s)
}

// HasChineseOrSymbols 是否有中文字符
func HasChineseOrSymbols(s string) bool {
	// 中文字符的正则表达式
	chineseCharRegex := regexp.MustCompile(`[\p{Han}]`) // 使用Unicode属性\p{Han}匹配任何中文字符

	// 中文符号的正则表达式（这里仅列举了一些常见的）
	chineseSymbolsRegex := regexp.MustCompile(`，。；：？！“”‘’（）《》【】、`)

	// 检查字符串是否包含中文字符或中文符号
	return chineseCharRegex.MatchString(s) || chineseSymbolsRegex.MatchString(s)
}

func ValidatePassword(password string) bool {
	// 定义密码的正则表达式，这里允许大小写字母、数字以及常见特殊字符
	// 注意：特殊字符的定义取决于你的需求，这里只是一个示例
	passwordRegex := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*(),.?":{}|<>]+$`)

	// 使用正则表达式进行匹配
	return passwordRegex.MatchString(password)
}
