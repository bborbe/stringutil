package stringutil

import (
	"strings"
)

func StringAfter(content string, find string) string {
	if len(find) == 0 {
		return content
	}
	pos := strings.Index(content, find)
	if pos == -1 {
		return ""
	}
	return content[pos + len(find):]
}

func StringBefore(content string, find string) string {
	if len(find) == 0 {
		return content
	}
	pos := strings.Index(content, find)
	if pos == -1 {
		return ""
	}
	return content[:pos]
}

func Trim(content string) string {
	runes := []rune(content)
	if len(runes) > 0 {
		if trimableChar(runes[0]) {
			return Trim(content[1:])
		}
		if trimableChar(runes[len(content) - 1]) {
			return Trim(content[:len(content) - 1])
		}
	}
	return content
}

func trimableChar(c rune) bool {
	return c == ' ' || c == '\n' || c == '\r'
}

func StringLess(a, b string) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return true
		}
		if a[i] > b[i] {
			return false
		}
	}
	return len(a) < len(b)
}
