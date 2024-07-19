package util

import (
	"github.com/lostvip-com/lv_framework/logme"
	"reflect"
)

/**
 *  翻译中文 到 它国语言 ，按 手动翻译，指定具体的field
 */
func TranslateField(dstPtr interface{}, keyField string) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			logme.Error(e)
		}
	}()

	dstValue1 := reflect.ValueOf(dstPtr)
	key := dstValue1.Elem().FieldByName(keyField).String()
	localeValue := GetTextLocale(key)
	dstValue1.Elem().FieldByName(keyField).SetString(localeValue)

	// dst必须结构体指针类型
	println(dstPtr)

	return nil
}

/**
 * 翻译中文 到 它国语言 ,遍历所有field.按tag中是否存在locale, 自动翻译
 */
func TranslateByTag(dstPtr interface{}) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			logme.Error(e)
		}
	}()

	dstType1 := reflect.TypeOf(dstPtr)
	dstValue1 := reflect.ValueOf(dstPtr)

	// dst必须结构体指针类型
	if dstType1.Kind() != reflect.Ptr {
		logme.Error("dst type should be a struct pointer : ", dstType1.Kind())
		panic("TranslateByTag --------dst type should be a struct pointer ")
	}
	if dstType1.Elem().Kind() != reflect.Struct {
		logme.Error("dstType1.Elem().Kind() should be a struct", dstType1.Elem().Kind())
		panic("CopyProperties XXXXXXXXXXXX--------dstType1.Elem().Kind() should be a struct  ：")
	}
	// 取具体内容
	elemType := dstType1.Elem()
	ememValue := dstValue1.Elem()
	// dst的属性个数
	dstPropertyNums := elemType.NumField()
	for i := 0; i < dstPropertyNums; i++ {
		fieldValue := ememValue.Field(i)
		// 无效，说明src没有这个属性 || 属性同名但类型不同
		if !fieldValue.IsValid() {
			continue
		}
		if !fieldValue.CanSet() {
			continue
		}
		fieldType := elemType.Field(i)
		tag := fieldType.Tag
		if tag == "" {
			continue
		}
		locale := tag.Get("locale")
		if locale != "auto" {
			continue
		}
		//这里修改 中文key 到它国语言
		key := fieldValue.String() //中文key
		localeValue := GetTextLocale(key)
		if err != nil {
			logme.Warn("XXXXXXXXXx=========>GetLocaleValueByKey:", fieldValue, locale, err)
			continue
		}
		fieldValue.Set(reflect.ValueOf(localeValue))

	}

	return nil
}

//
//
//
///**
// * 翻译中文 到 它国语言 ,遍历所有field.按tag中是否存在locale, 自动翻译
// */
//func TranslateSliceByTag(slicePtr []interface{}) (err error) {
//	// 防止意外panic
//	defer func() {
//		if e := recover(); e != nil {
//			logMy.Error(e)
//		}
//	}()
//	for ele := range slicePtr {
//		err = TranslateByTag(ele)
//		if err!=nil{
//			break
//		}
//	}
//
//	return err
//}
