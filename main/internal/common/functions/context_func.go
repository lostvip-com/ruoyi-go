package functions

import (
	"errors"
	"github.com/lostvip-com/lv_framework/logme"
	"github.com/lostvip-com/lv_framework/lv_cache/lv_ram"
	"github.com/lostvip-com/lv_framework/lv_global"
	"main/internal/common/global"
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
	ossUrl := GetValueFromRam(global.KeyOssUrl)
	return ossUrl
}

// 根据用户id和权限字符串判断是否有此权限
func AddInt(a, b int) int {
	return a + b
}

// 根据键获取值
func GetValueFromRam(key string) string {
	// 这里为了提高速度使用内在缓存
	result, err := lv_ram.GetRamCacheClient().Get(key)
	if err != nil {
		entity := &model.SysConfig{ConfigKey: key}
		err := entity.FindOne()
		if err != nil {
			panic(errors.New("获取配置失败,检查配置表：sys_config 中是否存在配置项：" + key))
		}
		result = entity.ConfigValue
		//内存续期
		lv_ram.GetRamCacheClient().Set(key, result, 10*time.Minute)
	}
	return result
}
