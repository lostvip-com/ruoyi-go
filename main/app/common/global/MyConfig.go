package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lostvip.com/cache/myredis"
	"lostvip.com/conf"
	"lostvip.com/logme"
	"os"
)

var cfg *MyConfig

func init() {
	GetConfigInstance()
}

type MyConfig struct {
	conf.ConfigDefault
}

func GetConfigInstance() conf.IConfig {
	if cfg == nil {
		cfg = &MyConfig{}
		cfg.LoadConf()
		conf.RegisterCfg(cfg)
		//日志
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
}

func resetCfg() {
	if conf.Config().IsDebug() {
		logme.Log.Warn(conf.Config().GetAppName() + " ============ gin debug模式,swagger 开启 ==========")
		gin.SetMode(gin.DebugMode)
		os.Setenv(conf.KEY_SWAGGER_OFF, "")
	} else {
		logme.Log.Warn(conf.Config().GetAppName() + " ============ gin 发布模式,swagger 禁用 ============")
		gin.SetMode(gin.ReleaseMode)
		os.Setenv(conf.KEY_SWAGGER_OFF, "off")
	}
	myredis.GetInstance()
	fmt.Println(conf.Config().GetAppName() + " ############# 数据库初始化完毕 ####################")
	InitTables()
	fmt.Println(conf.Config().GetAppName() + " ############# 修改表结构完毕 ####################")

}

func InitTables() {
	fmt.Println("------------ demo 自动建表 ------------")
	//setTableOption("App信息").AutoMigrate(model.BizApp{})
	//setTableOption("评论信息").AutoMigrate(model.AppComments{})
	//setTableOption("网友建议").AutoMigrate(model.AppSuggest{})
	//setTableOption("跳绳结果").AutoMigrate(model.BizMember{})
	////2023
	//setTableOption("跳绳结果").AutoMigrate(model.SptRope{})
	//setTableOption("50m短跑结果").AutoMigrate(model.SptRun{})
	//setTableOption("体适能结果").AutoMigrate(model.SptFitness{})
	//2022
	fmt.Println("------------ end 自动建表 ------------")
}
