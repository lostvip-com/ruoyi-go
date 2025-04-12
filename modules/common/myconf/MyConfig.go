package myconf

import (
	functions2 "common/functions"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_conf"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/utils/lv_logic"
	"html/template"
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
		//新的规则
		"Ctx":        functions2.GetCtx,
		"CtxPath":    functions2.CtxPath,
		"HasPerm":    functions2.HasPermi,
		"PermButton": functions2.PermButton,

		"DictSelect": functions2.DictSelect,
		"DictRadio":  functions2.DictRadio,
		"DictLabel":  functions2.DictLabel,
		"DictType":   functions2.DictType,

		"AddInt":    functions2.AddInt,
		"Contains":  functions2.Contains,
		"Copyright": functions2.GetCopyright,
		"OssUrl":    functions2.GetOssUrl,

		//兼容旧的
		"hasPermi":          functions2.HasPermi,          //兼容旧的
		"getDictTypeSelect": functions2.GetDictTypeSelect, //兼容旧的名称
		"getDictTypeRadio":  functions2.DictRadio,         //兼容旧的名称
		"contains":          functions2.Contains,          //兼容旧的名称
	}

	return mp
}

func (e *MyConfig) GetPartials() []string {
	return []string{"header", "footer", "system/menu/icon"}
} //
