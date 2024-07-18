package util

import (
	"encoding/json"
	"github.com/lostvip-com/lv_framework/logme"
	"io/ioutil"
)

var localeMap map[string]string

func LoadFileByLocale(locale string) error {
	var dataMap map[string]string
	if locale == "" || locale == "zh" {
		localeMap = nil
		return nil
	}
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	bytes, err := ioutil.ReadFile("locales/" + locale + ".json")
	if err != nil {
		return err
	}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(bytes, &dataMap)
	if err != nil {
		return err
	}

	localeMap = dataMap

	return err
}

func GetTextLocale(key string) string {
	if localeMap == nil {
		return key
	}
	v := localeMap[key]
	//读取的数据为json格式，需要进行解码
	if v == "" {
		v = key
		logme.Warn("XXX 语言文件中未找到对应的 value，使用默认值 " + key)
	}
	return v
}

func GetDataMap() map[string]string {
	return localeMap
}
