package functions

import (
	"context"
	"lostvip.com/cache/myredis"
	"lostvip.com/conf"
	"lostvip.com/logme"
	"lostvip.com/utils/lv_err"
	"main/app/system/model"
	"strings"

	"time"
)

func GetCopyright() string {
	copyright := conf.Config().GetConf("copyright")
	return copyright
}

func GetCtx() string {
	return conf.Config().GetContextPath()
}
func GetCtxPath(url string) string {
	ctxPath := conf.Config().GetContextPath()
	if !strings.HasPrefix(url, "http") {
		if strings.HasPrefix(url, "/") {
			url = ctxPath + url
		} else {
			url = ctxPath + "/" + url
		}
	} else {
		logme.Info("外链：" + ctxPath)
	}
	return url
}
func GetOssUrl() string {
	ossUrl := GetValueByKey("system.resource.url")
	if ossUrl == "" {
		ossUrl = "/static"
	} else if !strings.HasPrefix(ossUrl, "http") {
		ossUrl = conf.Config().GetContextPath() + ossUrl
	}
	return ossUrl
}

// 根据用户id和权限字符串判断是否有此权限
func AddInt(a, b int) int {
	return a + b
}

// 根据键获取值
func GetValueByKey(key string) string {
	//从缓存读取
	result := myredis.GetInstance().Get(context.Background(), key).Val()
	if result == "" {
		entity := &model.SysConfig{ConfigKey: key}
		err := entity.FindOne()
		lv_err.HasErrAndPanic(err)
		result = entity.ConfigValue
		myredis.GetInstance().SetEx(context.Background(), key, result, 10*time.Minute)
	}

	return result
}
