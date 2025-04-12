package lv_conf

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_file"
	"github.com/lostvip-com/lv_framework/utils/lv_net"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type ConfigDefault struct {
	vipperCfg     *viper.Viper
	proxyMap      map[string]string
	proxyEnable   bool
	cacheTpl      bool //默认不缓存模板，方便调试
	contextPath   string
	resourcesPath string
	logLevel      string
	autoMigrate   string
}

func (e ConfigDefault) GetResourcesPath() string {
	if e.resourcesPath == "" {
		e.resourcesPath = e.GetValueStr("go.application.resources-path")
	}
	return e.resourcesPath
}

func (e ConfigDefault) GetUploadPath() string {
	if e.resourcesPath == "" {
		e.resourcesPath = e.GetValueStr("go.application.upload-path")
	}
	return e.resourcesPath
}
func (e ConfigDefault) GetTmpPath() string {
	return "tmp" //固定临时文件目录
}
func (e ConfigDefault) GetGrpcPort() string {
	return e.GetValueStr("server.grpc.port")
}
func (e ConfigDefault) GetHost() string {
	return e.GetValueStr("server.host")
}
func (e *ConfigDefault) IsProxyEnabled() bool {
	return false
}

func (e *ConfigDefault) GetFuncMap() template.FuncMap {
	mp := template.FuncMap{}
	return mp
}

func (e *ConfigDefault) IsCacheTpl() bool {
	return e.cacheTpl
}

func (e *ConfigDefault) SetCacheTpl(cache bool) {
	e.cacheTpl = cache
}

func (e *ConfigDefault) GetVipperCfg() *viper.Viper {
	return e.vipperCfg
}

func (e *ConfigDefault) GetValueStrDefault(key string, defaultVal string) string {
	val := e.GetValueStr(key)
	if val == "" {
		val = defaultVal
	}
	return val
}

func (e *ConfigDefault) GetValueStr(key string) string {

	if e.vipperCfg == nil {
		e.LoadConf()
	}
	val := cast.ToString(e.vipperCfg.Get(key))
	if strings.HasPrefix(val, "$") { //存在动态表达式
		val = strings.TrimSpace(val)             //去空格
		val = lv_conv.SubStr(val, 2, len(val)-1) //去掉 ${}
		if strings.HasPrefix(val, "\"") {
			panic("${...} format error !!!")
		}
		index := strings.Index(val, ":") //ssz:按第一个: 分割，前半部分是占位符，后半部分是默认值
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
		e.LoadConf()
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

func (e *ConfigDefault) LoadConf() {
	e.vipperCfg = viper.New()
	if lv_file.IsFileExist("bootstrap.yml") || lv_file.IsFileExist("bootstrap.yaml") {
		e.vipperCfg.SetConfigName("bootstrap")
		e.vipperCfg.SetConfigType("yaml")
		e.vipperCfg.AddConfigPath("./")
		e.vipperCfg.ReadInConfig()
	} else {
		fmt.Println("未找到配置文件 bootstrap.yml 将使用默认配置！！！")
	}
	//加载第二个配置文件
	if lv_file.IsFileExist("application.yml") || lv_file.IsFileExist("application.yaml") {
		e.vipperCfg.SetConfigName("application")
		e.vipperCfg.SetConfigType("yaml")
		e.vipperCfg.AddConfigPath("./")
		e.vipperCfg.MergeInConfig()
	} else {
		fmt.Println("未找到配置文件 application.yml 将使用默认配置！！！")
	}
	if e.vipperCfg.GetBool("go.proxy.enable") == true {
		e.proxyEnable = true
		e.GetProxyMap()
	} else {
		fmt.Println("!!！！！！！！！！！！！！!!! porxy feature is disabled ！！！！！！！！！！！！！！！！！！！！！！！")
		e.proxyEnable = false
	}
}

/**
 * app port
 */
func (e *ConfigDefault) GetServerPort() int {
	port := e.GetValueStr("server.port")
	if port == "" {
		port = "8080"
	}
	return cast.ToInt(port)
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
	return e.contextPath
}

func (e *ConfigDefault) SetContextPath(ctxPath string) {
	e.contextPath = ctxPath
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
		master = "data.lv_db"
	}
	return master
}
func (e *ConfigDefault) GetSlave() string {
	return e.GetValueStr("go.datasource.slave")
}

// IsDebug todo
func (e *ConfigDefault) GetLogLevel() string {
	if e.logLevel == "" {
		e.logLevel = e.GetValueStr("go.log.level")
	}
	return e.logLevel
}

func (e *ConfigDefault) GetAutoMigrate() string {
	if e.autoMigrate == "" {
		e.autoMigrate = e.GetValueStr("go.datasource.auto-migrate")
	}
	return e.autoMigrate
}

func (e *ConfigDefault) GetLogOutput() string {
	output := e.GetValueStr("go.log.output")
	return output
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

func (e *ConfigDefault) IsProxyEnable() bool {
	return e.proxyEnable
}

func (e *ConfigDefault) GetProxyMap() *map[string]string {
	if e.proxyEnable && e.proxyMap == nil {
		e.LoadProxyInfo()
	}
	return &e.proxyMap
}

func (e *ConfigDefault) LoadProxyInfo() *map[string]string {
	fmt.Println("######### 加载代理配置信息 start #############")
	if !e.IsProxyEnable() {
		return nil
	}
	list := e.GetVipperCfg().GetStringSlice("go.proxy.prefix")
	e.proxyMap = make(map[string]string)
	for _, v := range list {
		index := strings.Index(v, "=")
		key := lv_conv.SubStr(v, 0, index)
		hostPort := lv_conv.SubStr(v, index+1, len(v))
		e.proxyMap[key] = hostPort
	}
	e.proxyEnable = e.GetBool("go.proxy.enable")
	fmt.Println("go.proxy:", e.proxyMap)
	fmt.Println("######### 加载代理配置信息 end #############")
	return &e.proxyMap
}

func (e *ConfigDefault) GetPartials() []string {
	return []string{}
} //
