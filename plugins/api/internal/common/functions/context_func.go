package functions

import (
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/lv_log"
	"strings"
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
		lv_log.Info("外链：" + ctxPath)
	}
	return url
}

func GetOssUrl() string {
	ossUrl := "/static"
	return ossUrl
}

// 根据用户id和权限字符串判断是否有此权限
func AddInt(a, b int) int {
	return a + b
}
