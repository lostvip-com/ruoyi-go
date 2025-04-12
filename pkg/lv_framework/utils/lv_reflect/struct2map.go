package lv_reflect

import (
	"fmt"
	"reflect"
)

func CopyProperties2Map(input interface{}, result map[string]interface{}) error {
	// 检查输入是否为指针或结构体类型
	value := reflect.ValueOf(input)
	if value.Kind() == reflect.Ptr {
		value = value.Elem() // 解引用指针
	}

	if value.Kind() != reflect.Struct {
		return fmt.Errorf("input must be a struct or a pointer to a struct")
	}

	typeOfStruct := value.Type()

	// 遍历结构体字段
	for i := 0; i < value.NumField(); i++ {
		field := typeOfStruct.Field(i)
		fieldName := field.Name
		fieldValue := value.Field(i)

		// 只添加可导出的字段（首字母大写）
		if fieldValue.CanInterface() {
			result[fieldName] = fieldValue.Interface()
		}
	}

	return nil
}
