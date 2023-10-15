package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

var iconfig IConfig

func Config() IConfig {
	if iconfig == nil {
		panic("implement not found for interface IConfig, forgot register?")
	}
	return iconfig
}

func RegisterCfg(iconf IConfig) {
	fmt.Println("=============lib_framework=======RegisterConf===========")
	iconfig = iconf
}

type IConfig interface {
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
	LoadConf()
	IsDebug() bool
	GetVipperCfg() *viper.Viper
	GetConf(key string) string
}
