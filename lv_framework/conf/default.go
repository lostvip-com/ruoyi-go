package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"lostvip.com/utils/lib_file"
	"lostvip.com/utils/lib_net"
)

type ConfigDefault struct {
	VipperCfg *viper.Viper
}

func (e *ConfigDefault) GetVipperCfg() *viper.Viper {
	return e.VipperCfg
}

func (e *ConfigDefault) LoadConf() {
	e.VipperCfg = viper.New()
	if lib_file.FileExist("bootstrap.yml") || lib_file.FileExist("bootstrap.yaml") {
		e.VipperCfg.SetConfigName("bootstrap")
		e.VipperCfg.SetConfigType("yaml")
		e.VipperCfg.AddConfigPath("./")
		e.VipperCfg.ReadInConfig()
	}
	//加载第二个配置文件
	if lib_file.FileExist("application.yml") || lib_file.FileExist("application.yaml") {
		e.VipperCfg.SetConfigName("application")
		e.VipperCfg.SetConfigType("yaml")
		e.VipperCfg.AddConfigPath("./")
		e.VipperCfg.MergeInConfig()
	}
}

func (e *ConfigDefault) GetPort() int {
	Port := e.VipperCfg.GetInt("server.port")
	if Port == 0 {
		Port = 8080
	}
	return Port
}

/**
 * app port
 */
func (e *ConfigDefault) GetServerPort() int {
	port := e.VipperCfg.GetInt("server.port")
	if port == 0 {
		port = 8080
	}
	return port
}

/**
 * app port
 */
func (e *ConfigDefault) GetServerIP() string {
	ip := e.VipperCfg.GetString("server.ip")
	if ip == "" {
		ip = lib_net.GetLocaHost()
	}
	return ip
}

func (e *ConfigDefault) GetContextPath() string {
	path := e.VipperCfg.GetString("server.context-path")
	return path
}

func (e *ConfigDefault) GetConf(key string) string {
	v := e.VipperCfg.GetString(key)
	return v
}

var appName string

func (e *ConfigDefault) GetAppName() string {
	if appName == "" {
		appName = e.VipperCfg.GetString("go.application.name")
		if appName == "" {
			appName = "whoami"
		}
	}
	return appName
}
func (e *ConfigDefault) GetDriver() string {
	driver := e.VipperCfg.GetString("go.datasource.driver")
	if driver == "" {
		driver = "sqlite3"
	}
	return driver
}
func (e *ConfigDefault) GetMaster() string {
	master := e.VipperCfg.GetString("go.datasource.master")
	if master == "" {
		master = "data.db"
	}
	return master
}
func (e *ConfigDefault) GetSlave() string {
	return e.VipperCfg.GetString("go.datasource.slave")
}

func (e *ConfigDefault) IsDebug() bool {
	debug := e.VipperCfg.GetBool("go.application.debug")
	return debug
}

func (e *ConfigDefault) GetAppActive() string {
	return e.VipperCfg.GetString("go.application.active")
}

func (e *ConfigDefault) GetNacosAddrs() string {
	return e.VipperCfg.GetString("go.cloud.nacos.discovery.server-addr")
}

func (e *ConfigDefault) GetNacosPort() int {
	port := e.VipperCfg.GetInt("go.cloud.nacos.discovery.port")
	if port == 0 {
		port = 8848
	}
	return port
}
func (e *ConfigDefault) GetNacosNamespace() string {
	ns := e.VipperCfg.GetString("go.cloud.nacos.discovery.namespace")
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
