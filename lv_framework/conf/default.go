package conf

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_file"
	"lostvip.com/utils/lv_net"
	"os"
	"strings"
)

type ConfigDefault struct {
	vipperCfg *viper.Viper
}

func (e *ConfigDefault) GetVipperCfg() *viper.Viper {
	return e.vipperCfg
}

func (e *ConfigDefault) GetValueStr(key string) string {

	if e.vipperCfg == nil {
		e.vipperCfg = e.LoadConf()
	}
	val := cast.ToString(e.vipperCfg.Get(key))
	if strings.HasPrefix(val, "$") { //存在动态表达式
		val = strings.TrimSpace(val)             //去空格
		val = lv_conv.SubStr(val, 2, len(val)-1) //去掉 ${}
		index := strings.Index(val, ":")         //ssz:按第一个: 分割，前半部分是占位符，后半部分是默认值
		val0 := lv_conv.SubStr(val, 0, index)
		val0 = os.Getenv(val0) //从环境变量中取值,替换
		if val0 == "" {        //未设置环境变量,使用默认值
			val = lv_conv.SubStr(val, index+1, len(val))
			val = strings.Trim(val, "\"")
		} else {
			val = val0
		}
	}
	return val
}

func (e *ConfigDefault) GetBool(key string) bool {
	if e.vipperCfg == nil {
		e.vipperCfg = e.LoadConf()
	}
	val := cast.ToString(e.vipperCfg.Get(key))
	val = e.parseVal(val)
	if val == "true" {
		return true
	} else {
		return false
	}
}

func (e *ConfigDefault) parseVal(val string) string {
	if strings.HasPrefix(val, "$") { //存在动态表达式
		val = strings.TrimSpace(val)             //去空格
		val = lv_conv.SubStr(val, 2, len(val)-1) //去掉 ${}
		index := strings.Index(val, ":")         //ssz:按第一个: 分割，前半部分是占位符，后半部分是默认值
		val0 := lv_conv.SubStr(val, 0, index)
		val0 = os.Getenv(val0) //从环境变量中取值,替换
		if val0 == "" {        //未设置环境变量,使用默认值
			val = lv_conv.SubStr(val, index, len(val))
		} else {
			val = val0
		}
	}
	return val
}

func (e *ConfigDefault) LoadConf() *viper.Viper {
	e.vipperCfg = viper.New()
	if lv_file.IsFileExist("bootstrap.yml") || lv_file.IsFileExist("bootstrap.yaml") {
		e.vipperCfg.SetConfigName("bootstrap")
		e.vipperCfg.SetConfigType("yaml")
		e.vipperCfg.AddConfigPath("./")
		e.vipperCfg.ReadInConfig()
	}
	//加载第二个配置文件
	if lv_file.IsFileExist("application.yml") || lv_file.IsFileExist("application.yaml") {
		e.vipperCfg.SetConfigName("application")
		e.vipperCfg.SetConfigType("yaml")
		e.vipperCfg.AddConfigPath("./")
		e.vipperCfg.MergeInConfig()
	}
	return e.vipperCfg
}

func (e *ConfigDefault) GetPort() int {
	Port := e.vipperCfg.GetInt("server.port")
	if Port == 0 {
		Port = 8080
	}
	return Port
}

/**
 * app port
 */
func (e *ConfigDefault) GetServerPort() int {
	port := e.vipperCfg.GetInt("server.port")
	if port == 0 {
		port = 8080
	}
	return port
}

/**
 * app port
 */
func (e *ConfigDefault) GetServerIP() string {
	ip := e.GetValueStr("server.ip")
	if ip == "" {
		ip = lv_net.GetLocaHost()
	}
	return ip
}

func (e *ConfigDefault) GetContextPath() string {
	path := e.GetValueStr("server.context-path")
	return path
}

func (e *ConfigDefault) GetConf(key string) string {
	v := e.GetValueStr(key)
	return v
}

var appName string

func (e *ConfigDefault) GetAppName() string {
	if appName == "" {
		appName = e.GetValueStr("go.application.name")
		if appName == "" {
			appName = "lostvip"
		}
	}
	return appName
}
func (e *ConfigDefault) GetDriver() string {
	driver := e.GetValueStr("go.datasource.driver")
	if driver == "" {
		driver = "sqlite3"
	}
	return driver
}
func (e *ConfigDefault) GetMaster() string {
	master := e.GetValueStr("go.datasource.master")
	if master == "" {
		master = "data.db"
	}
	return master
}
func (e *ConfigDefault) GetSlave() string {
	return e.GetValueStr("go.datasource.slave")
}

func (e *ConfigDefault) IsDebug() bool {
	debug := e.GetBool("go.application.debug")
	return debug
}

func (e *ConfigDefault) GetAppActive() string {
	return e.GetValueStr("go.application.active")
}

func (e *ConfigDefault) GetNacosAddrs() string {
	return e.GetValueStr("go.cloud.nacos.discovery.server-addr")
}

func (e *ConfigDefault) GetNacosPort() int {
	port := e.vipperCfg.GetInt("go.cloud.nacos.discovery.port")
	if port == 0 {
		port = 8848
	}
	return port
}
func (e *ConfigDefault) GetNacosNamespace() string {
	ns := e.GetValueStr("go.cloud.nacos.discovery.namespace")
	return ns
}
func (e *ConfigDefault) GetGroupDefault() string {
	return "DEFAULT_GROUP"
}
func (e *ConfigDefault) GetDataId() string {
	key := e.GetAppName() + "-" + e.GetAppActive() + ".yml"
	fmt.Println(" dataId: " + key)
	return key
}
