package server

import (
	"regexp"
	"strings"
)

// ConvertLE is used to normaliza line endings when exchanging clipboard content.
func ConvertLE(text, op string) string {
	switch {
	case strings.EqualFold("lf", op):
		text = strings.Replace(text, "\r\n", "\n", -1)
		return strings.Replace(text, "\r", "\n", -1)
	case strings.EqualFold("crlf", op):
		text = regexp.MustCompile(`\r(.)|\r$`).ReplaceAllString(text, "\r\n$1")
		text = regexp.MustCompile(`([^\r])\n|^\n`).ReplaceAllString(text, "$1\r\n")
		return text
	default:
		return text
	}
}
