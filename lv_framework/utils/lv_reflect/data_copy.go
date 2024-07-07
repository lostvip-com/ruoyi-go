package lv_reflect

import (
	"fmt"
	"reflect"
)

// CopyStruct
// dst 目标结构体，src 源结构体
// 必须传入指针，且不能为nil
// 它会把src与dst的相同字段名的值，复制到dst中
//func SimpleCopyProperties(src, dst interface{}, copyNil bool) (err error) {
//	// 防止意外panic
//	defer func() {
//		if e := recover(); e != nil {
//			err = errors.New(fmt.Sprintf("%v", e))
//		}
//	}()
//
//	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
//	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)
//
//	// dst必须结构体指针类型
//	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
//		return errors.New("dst type should be a struct pointer")
//	}
//
//	// src必须为结构体或者结构体指针
//	if srcType.Kind() == reflect.Ptr {
//		srcType, srcValue = srcType.Elem(), srcValue.Elem()
//	}
//	if srcType.Kind() != reflect.Struct {
//		return errors.New("src type should be a struct or a struct pointer")
//	}
//
//	// 取具体内容
//	dstType, dstValue = dstType.Elem(), dstValue.Elem()
//
//	// 属性个数
//	propertyNums := dstType.NumField()
//
//	for i := 0; i < propertyNums; i++ {
//		// 属性
//		property := dstType.Field(i)
//		// 待填充属性值
//		propertyValue := srcValue.FieldByName(property.Name)
//
//		// 无效，说明src没有这个属性 || 属性同名但类型不同
//		if !propertyValue.IsValid() || property.Type != propertyValue.Type() {
//			continue
//		}
//
//		if dstValue.Field(i).CanSet() {
//			if copyNil {
//				dstValue.Field(i).Set(propertyValue)
//			} else {
//				if propertyValue=={
//					dstValue.Field(i).Set(propertyValue)
//				}
//
//			}
//
//		}
//	}
//
//	return nil
//}

/**
 * ai 生成,只处理非空字段
 */
func CopyProperties(source interface{}, target interface{}) error {
	sourceValue := reflect.Indirect(reflect.ValueOf(source))
	targetValue := reflect.Indirect(reflect.ValueOf(target))

	if sourceValue.Kind() != reflect.Struct || targetValue.Kind() != reflect.Ptr || !targetValue.Elem().CanSet() {
		return fmt.Errorf("invalid arguments")
	}

	for i := 0; i < sourceValue.NumField(); i++ {
		fieldType := sourceValue.Type().Field(i).Type
		value := sourceValue.Field(i).Interface()

		// 只处理非空字段
		if value == nil && fieldType.Kind() != reflect.Slice && fieldType.Kind() != reflect.Map {
			continue
		}

		targetField := targetValue.Elem().FieldByIndex([]int{i})
		if !targetField.IsValid() || !targetField.CanSet() {
			continue
		}

		switch fieldType.Kind() {
		case reflect.String:
			fallthrough //强制执行下面的case
		case reflect.Int:
			fallthrough //强制执行下面的case
		case reflect.Bool:
			targetField.Set(reflect.ValueOf(value))
		default:
			err := CopyProperties(value, targetField.Addr().Interface())
			if err != nil {
				return err
			}
		}
	}

	return nil
}
func IsMap(data interface{}) bool {
	t := reflect.TypeOf(data)
	return t.Kind() == reflect.Map
}
