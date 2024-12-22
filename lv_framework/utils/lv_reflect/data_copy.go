package lv_reflect

import (
	"github.com/jinzhu/copier"
	"reflect"
)

func CopyProperties(fromValue interface{}, toValue interface{}) error {
	err := copier.CopyWithOption(toValue, fromValue, copier.Option{IgnoreEmpty: true})
	return err
}

func CopyProp(fromValue interface{}, toValue interface{}, ignoreEmpty bool) error {
	//dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	//srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)
	//// dst必须结构体指针类型
	//if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
	//	return errors.New("dst type should be a struct pointer")
	//}
	err := copier.CopyWithOption(toValue, fromValue, copier.Option{IgnoreEmpty: ignoreEmpty})
	return err
}

func IsMap(data interface{}) bool {
	t := reflect.TypeOf(data)
	return t.Kind() == reflect.Map
}

func Map2Struct(sourceMap map[any]any, destPtr any) error {
	err := copier.Copy(destPtr, sourceMap)
	return err
}
