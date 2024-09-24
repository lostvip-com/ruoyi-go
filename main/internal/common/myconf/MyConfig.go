package myconf

import (
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_conf"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/utils/lv_logic"
	"html/template"
	"main/internal/common/functions"
	"os"
)

var cfg *MyConfig

func init() {
	GetConfigInstance()
}

type MyConfig struct {
	lv_conf.ConfigDefault
}

func (e *MyConfig) IsProxyEnabled() bool {
	return e.IsProxyEnable()
}

func GetConfigInstance() lv_global.IConfig {
	if cfg == nil {
		cfg = &MyConfig{}
		cfg.LoadConf()
		lv_global.RegisterCfg(cfg)
	}
	return cfg
}

// LoadConf 预加载部分参数，避免每次判断
func (e *MyConfig) LoadConf() {
	e.ConfigDefault.LoadConf() //调用
	e.SetCacheTpl(e.GetBool("go.application.cache-tpl"))
	path := e.GetValueStr("server.context-path")
	e.SetContextPath(path)
	////远程
	//nacosAddrs := e.GetNacosAddrs()
	//if nacosAddrs != "" { //加载远程配置文件
	//	nacos.LoadRemoteConfig(e.ConfigDefault, resetCfg)
	//}
	//日志
	lv_global.IsDebug = lv_logic.IfTrue(e.GetLogLevel() == "debug", true, false).(bool)
	lv_global.LogOutputType = e.GetLogOutput()
	return
}

func resetCfg() {
	conig := lv_global.Config()
	if lv_global.IsDebug {
		lv_log.Warn(conig.GetAppName() + " ============ gin debug模式,swagger 开启 ==========")
		gin.SetMode(gin.DebugMode)
		os.Setenv(lv_global.KEY_SWAGGER_OFF, "")
	} else {
		lv_log.Warn(conig.GetAppName() + " ============ gin 发布模式,swagger 禁用 ============")
		gin.SetMode(gin.ReleaseMode)
		os.Setenv(lv_global.KEY_SWAGGER_OFF, "off")
	}
	if conig.GetLogLevel() == "debug" {
		lv_log.Warn(conig.GetAppName() + " ============ gin debug模式,swagger 开启 ==========")
		gin.SetMode(gin.DebugMode)
		os.Setenv(lv_global.KEY_SWAGGER_OFF, "")
	} else {
		lv_log.Warn(conig.GetAppName() + " ============ gin 发布模式,swagger 禁用 ============")
		gin.SetMode(gin.ReleaseMode)
		os.Setenv(lv_global.KEY_SWAGGER_OFF, "off")
	}
}

func (e *MyConfig) GetFuncMap() template.FuncMap {

	mp := template.FuncMap{
		"hasPermi":          functions.HasPermi,
		"getPermiButton":    functions.GetPermiButton,
		"getDictLabel":      functions.GetDictLabel,
		"getDictTypeSelect": functions.GetDictTypeSelect,
		"getDictTypeRadio":  functions.GetDictTypeRadio,
		"getDictTypeData":   functions.GetDictTypeData,
		"Copyright":         functions.GetCopyright,
		"OssUrl":            functions.GetOssUrl,
		"Ctx":               functions.GetCtx,
		"getCtxPath":        functions.GetCtxPath,
		"addInt":            functions.AddInt,
		"contains":          functions.Contains,
	}

	return mp
}

func (e *MyConfig) GetPartials() []string {
	return []string{"header", "footer", "system/menu/icon"}
} //
