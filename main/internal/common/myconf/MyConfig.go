package myconf

import (
	"fmt"
	"github.com/alecthomas/template"
	"github.com/gin-gonic/gin"
	"github.com/lv_framework/conf"
	"github.com/lv_framework/logme"
	"lostvip.com/cache/lv_redis"
	"lostvip.com/lv_global"
	"main/internal/common/functions"
	"os"
)

var cfg *MyConfig

func init() {
	GetConfigInstance()
}

type MyConfig struct {
	conf.ConfigDefault
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

func (e *MyConfig) LoadConf() {
	e.ConfigDefault.LoadConf() //调用
	////远程
	//nacosAddrs := e.GetNacosAddrs()
	//if nacosAddrs != "" { //加载远程配置文件
	//	nacos.LoadRemoteConfig(e.ConfigDefault, resetCfg)
	//}
	//日志
	lv_global.IsDebug = e.IsDebug()
	lv_global.LogOutputType = e.GetLogOutput()
}

func resetCfg() {
	conig := lv_global.Config()
	if lv_global.IsDebug {
		logme.Warn(conig.GetAppName() + " ============ gin debug模式,swagger 开启 ==========")
		gin.SetMode(gin.DebugMode)
		os.Setenv(lv_global.KEY_SWAGGER_OFF, "")
	} else {
		logme.Warn(conig.GetAppName() + " ============ gin 发布模式,swagger 禁用 ============")
		gin.SetMode(gin.ReleaseMode)
		os.Setenv(lv_global.KEY_SWAGGER_OFF, "off")
	}
	if conig.IsDebug() {
		logme.Warn(conig.GetAppName() + " ============ gin debug模式,swagger 开启 ==========")
		gin.SetMode(gin.DebugMode)
		os.Setenv(lv_global.KEY_SWAGGER_OFF, "")
	} else {
		logme.Warn(conig.GetAppName() + " ============ gin 发布模式,swagger 禁用 ============")
		gin.SetMode(gin.ReleaseMode)
		os.Setenv(lv_global.KEY_SWAGGER_OFF, "off")
	}
	lv_redis.GetInstance()
	fmt.Println(conig.GetAppName() + " ############# 数据库初始化完毕 ####################")
	InitTables()
	fmt.Println(conig.GetAppName() + " ############# 修改表结构完毕 ####################")

}

func InitTables() {
	fmt.Println("------------ demo 自动建表 ------------")
	//setTableOption("App信息").AutoMigrate(model_cmn.BizApp{})
	//setTableOption("评论信息").AutoMigrate(model_cmn.AppComments{})
	//setTableOption("网友建议").AutoMigrate(model_cmn.AppSuggest{})
	//setTableOption("跳绳结果").AutoMigrate(model_cmn.BizMember{})
	////2023
	//setTableOption("跳绳结果").AutoMigrate(model_cmn.SptRope{})
	//setTableOption("50m短跑结果").AutoMigrate(model_cmn.SptRun{})
	//setTableOption("体适能结果").AutoMigrate(model_cmn.SptFitness{})
	//2022
	fmt.Println("------------ end 自动建表 ------------")
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
