package lv_global

import (
	"fmt"
	"html/template"

	"github.com/spf13/viper"
)

var iconfig IConfig

func Config() IConfig {
	return iconfig
}
func RegisterCfg(iconf IConfig) {
	fmt.Println("=============lib_framework=======RegisterConf===========")
	iconfig = iconf
}

type IConfig interface {
	GetUploadPath() string //用于提供对外服务地址
	GetResourcesPath() string
	GetTmpPath() string
	GetServerPort() int
	GetServerIP() string
	GetContextPath() string
	GetAppName() string
	GetDriver() string
	GetMaster() string
	GetSlave() string
	GetAppActive() string
	GetNacosAddrs() string
	GetNacosPort() int
	GetNacosNamespace() string
	GetGroupDefault() string

	GetDataId() string
	GetLogLevel() string
	IsCacheTpl() bool
	GetVipperCfg() *viper.Viper
	GetConf(key string) string
	GetValueStr(key string) string
	GetBool(key string) bool
	GetProxyMap() *map[string]string
	IsProxyEnabled() bool
	LoadConf()
	GetFuncMap() template.FuncMap
	GetAutoMigrate() string
	GetPartials() []string
	GetGrpcPort() string
	GetHost() string
}
