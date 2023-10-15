package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"lostvip.com/utils/lib_file"
	"lostvip.com/utils/lib_net"
)

type ConfigDefault struct {
}

var vipperCfg *viper.Viper

func (e *ConfigDefault) GetVipperCfg() *viper.Viper {
	return vipperCfg
}

func (e *ConfigDefault) LoadConf() {
	vipperCfg = viper.New()
	if lib_file.FileExist("bootstrap.yml") || lib_file.FileExist("bootstrap.yaml") {
		vipperCfg.SetConfigName("bootstrap")
		vipperCfg.SetConfigType("yaml")
		vipperCfg.AddConfigPath("./")
		vipperCfg.ReadInConfig()
	}
	//加载第二个配置文件
	if lib_file.FileExist("application.yml") || lib_file.FileExist("application.yaml") {
		vipperCfg.SetConfigName("application")
		vipperCfg.SetConfigType("yaml")
		vipperCfg.AddConfigPath("./")
		vipperCfg.MergeInConfig()
	}
}

func (e *ConfigDefault) GetPort() int {
	Port := vipperCfg.GetInt("server.port")
	if Port == 0 {
		Port = 8080
	}
	return Port
}

/**
 * app port
 */
func (e *ConfigDefault) GetServerPort() int {
	port := vipperCfg.GetInt("server.port")
	if port == 0 {
		port = 8080
	}
	return port
}

/**
 * app port
 */
func (e *ConfigDefault) GetServerIP() string {
	ip := vipperCfg.GetString("server.ip")
	if ip == "" {
		ip = lib_net.GetLocaHost()
	}
	return ip
}

func (e *ConfigDefault) GetContextPath() string {
	path := vipperCfg.GetString("server.context-path")
	return path
}

func (e *ConfigDefault) GetConf(key string) string {
	v := vipperCfg.GetString(key)
	return v
}

var appName string

func (e *ConfigDefault) GetAppName() string {
	if appName == "" {
		appName = vipperCfg.GetString("go.application.name")
		if appName == "" {
			appName = "whoami"
		}
	}
	return appName
}
func (e *ConfigDefault) GetDriver() string {
	driver := vipperCfg.GetString("go.datasource.driver")
	if driver == "" {
		driver = "sqlite3"
	}
	return driver
}
func (e *ConfigDefault) GetMaster() string {
	master := vipperCfg.GetString("go.datasource.master")
	if master == "" {
		master = "data.db"
	}
	return master
}
func (e *ConfigDefault) GetSlave() string {
	return vipperCfg.GetString("go.datasource.slave")
}

var debug bool

func (e *ConfigDefault) IsDebug() bool {
	if debug {
		debug = vipperCfg.GetBool("go.application.debug")
	}
	return debug
}

func (e *ConfigDefault) GetAppActive() string {
	return vipperCfg.GetString("go.application.active")
}

func (e *ConfigDefault) GetNacosAddrs() string {
	return vipperCfg.GetString("go.cloud.nacos.discovery.server-addr")
}

func (e *ConfigDefault) GetNacosPort() int {
	port := vipperCfg.GetInt("go.cloud.nacos.discovery.port")
	if port == 0 {
		port = 8848
	}
	return port
}
func (e *ConfigDefault) GetNacosNamespace() string {
	ns := vipperCfg.GetString("go.cloud.nacos.discovery.namespace")
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
