package lv_db

import (
	"strings"
)

/**
 * 首字母小写
 */
func ToCamelFirstLower(colName string) string {
	var fieldName string
	arrCol := strings.Split(colName, "_")
	for i := 0; i < len(arrCol); i++ {
		if arrCol[i] == "" {
			continue
		}
		arr := []byte(arrCol[i])
		strStart := string(arr[:1])
		strend := string(arr[1:])
		if i == 0 {
			fieldName = strStart + strend
		} else {
			fieldName = fieldName + strings.ToUpper(strStart) + strend
		}
	}

	return fieldName
}

/**
 * 首字母大写
 */
func ToCamelFirstUpper(colName string) string {
	var fieldName = ""
	arrCol := strings.Split(colName, "_")
	for i := 0; i < len(arrCol); i++ {
		if arrCol[i] == "" {
			continue
		}
		arr := []byte(arrCol[i])
		strStart := string(arr[:1])
		strend := string(arr[1:])

		fieldName = fieldName + strings.ToUpper(strStart) + strend
	}

	return fieldName
}
