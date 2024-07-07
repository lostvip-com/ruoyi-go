package lv_conv

import (
	"strings"
)

func SubStr(str string, startIndex, endIndex int) string {
	rs := []rune(str)
	return string(rs[startIndex:endIndex])
}

// 将带分隔符的字符串切成int64数组
func ToInt64Array(str, split string) []int64 {
	result := make([]int64, 0)
	if str == "" {
		return result
	}

	arr := strings.Split(str, split)

	if len(arr) > 0 {
		for i := range arr {
			if arr[i] != "" {
				result = append(result, Int64(arr[i]))
			}
		}
	}

	return result

}

// 过滤收尾有分隔符的数字字符串
func ReplaceHeadAndEndStr(str, split string) string {
	result := ""
	arr := strings.Split(str, split)
	if len(arr) <= 0 {
		return result
	}

	for i := range arr {
		if arr[i] != "" {
			if i == 0 {
				result = arr[i]
			} else {
				result += "," + arr[i]
			}
		}
	}

	return result
}
