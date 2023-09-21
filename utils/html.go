package utils

import (
	"github.com/kennygrant/sanitize"
	"regexp"
)

// GetHtmlRawContent get raw content
func GetHtmlRawContent(html string) string {
	var text = sanitize.HTML(html)
	//移除多余空白行
	regex, err := regexp.Compile("\n\n")
	if err != nil {
		return html
	}
	return regex.ReplaceAllString(text, "")
}
