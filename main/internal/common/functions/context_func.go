package functions

import (
	"context"
	"github.com/lostvip-com/lv_framework/cache/lv_redis"
	"github.com/lostvip-com/lv_framework/logme"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"main/internal/system/model"
	"strings"

	"time"
)

func GetCopyright() string {
	copyright := lv_global.Config().GetConf("copyright")
	return copyright
}

func GetCtx() string {
	return lv_global.Config().GetContextPath()
}
func GetCtxPath(url string) string {
	ctxPath := lv_global.Config().GetContextPath()
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
		ossUrl = lv_global.Config().GetContextPath() + ossUrl
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
	result := lv_redis.GetInstance().Get(context.Background(), key).Val()
	if result == "" {
		entity := &model.SysConfig{ConfigKey: key}
		err := entity.FindOne()
		lv_err.HasErrAndPanic(err)
		result = entity.ConfigValue
		lv_redis.GetInstance().SetEx(context.Background(), key, result, 10*time.Minute)
	}

	return result
}
