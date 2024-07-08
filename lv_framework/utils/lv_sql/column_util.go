package lv_sql

import (
	"strings"
	"unicode"
)

func ToTitle(column string) string {
	parts := strings.Split(column, "_")
	for i, part := range parts {
		parts[i] = FirstToUpper(part)
	}
	return strings.Join(parts, "")
}

func ToCamel(column string) string {
	parts := strings.Split(column, "_")
	for i, part := range parts {
		if i > 0 {
			parts[i] = FirstToUpper(part)
		}
	}
	return strings.Join(parts, "")
}

func FirstToUpper(str string) string {
	// 将字符串的第一个字符转换为大写字母
	firstChar := str[0]
	firstCharUpper := string(unicode.ToUpper(rune(firstChar)))
	str = firstCharUpper + str[1:]
	return str
}
