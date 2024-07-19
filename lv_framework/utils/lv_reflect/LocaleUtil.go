package util

import (
	"encoding/json"
	"github.com/lostvip-com/lv_framework/logme"
	"github.com/lostvip-com/lv_framework/lv_cache"
	"os"
	"time"
)

func LoadFileByLocale(locale string) error {
	var dataMap map[string]any
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	bytes, err := os.ReadFile("locales/" + locale + ".json")
	if err != nil {
		return err
	}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(bytes, &dataMap)
	if err != nil {
		return err
	}
	lv_cache.GetCacheClient().HMSet("locale", dataMap, time.Hour)
	return err
}

func GetTextLocale(key string) string {
	v, _ := lv_cache.GetCacheClient().HGet("locale", key)
	//读取的数据为json格式，需要进行解码
	if v == "" {
		v = key
		logme.Warn("XXX 语言文件中未找到对应的 value，使用默认值 " + key)
	}
	return v
}
