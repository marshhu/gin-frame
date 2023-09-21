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
