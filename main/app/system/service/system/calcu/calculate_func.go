package calcu

import "strings"

// 根据用户id和权限字符串判断是否有此权限
func AddInt(a, b int) int {
	return a + b
}

func Contains(str, subStr string) bool {
	return strings.Contains(str, subStr)
}
