package lv_reflect

import (
	"github.com/lostvip-com/lv_framework/lv_log"
	"reflect"
)

func TranslateField(dstPtr interface{}, keyField string, local string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			lv_log.Error(e)
		}
	}()
	dstValue1 := reflect.ValueOf(dstPtr)
	key := dstValue1.Elem().FieldByName(keyField).String()
	localeValue := GetTextLocale(local, key)
	dstValue1.Elem().FieldByName(keyField).SetString(localeValue)
	// dst必须结构体指针类型
	return
}

func TranslateByTag(dstPtr interface{}) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			lv_log.Error(e)
		}
	}()
	dstType1 := reflect.TypeOf(dstPtr)
	dstValue1 := reflect.ValueOf(dstPtr)
	// dst必须结构体指针类型
	if dstType1.Kind() != reflect.Ptr {
		lv_log.Error("dst type should be a struct pointer : ", dstType1.Kind())
		panic("TranslateByTag --------dst type should be a struct pointer ")
	}
	if dstType1.Elem().Kind() != reflect.Struct {
		lv_log.Error("dstType1.Elem().Kind() should be a struct", dstType1.Elem().Kind())
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
		localeValue := GetTextLocale(locale, key)
		if err != nil {
			lv_log.Warn("XXXXXXXXXx=========>GetLocaleValueByKey:", fieldValue, locale, err)
			continue
		}
		fieldValue.Set(reflect.ValueOf(localeValue))
	}
	return nil
}
