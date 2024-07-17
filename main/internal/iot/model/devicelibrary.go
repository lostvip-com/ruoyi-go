package models

import (
	"database/sql/driver"
	"encoding/json"
)

// 驱动库，驱动市场
type DeviceLibrary struct {
	Timestamps      `gorm:"embedded"`
	Id              string          `gorm:"id;primaryKey;not null;type:string;size:255;comment:主键"`
	Name            string          `gorm:"type:string;size:255;comment:名字"`
	Description     string          `gorm:"type:text;comment:描述"`
	Protocol        string          `gorm:"type:string;size:255;comment:驱动协议"`
	Version         string          `gorm:"type:string;size:255;comment:当前安装版本"`
	ContainerName   string          `gorm:"type:string;size:255;comment:容器名字"`
	DockerConfigId  string          `gorm:"type:string;size:255;comment:镜像仓库配置表id"`
	DockerRepoName  string          `gorm:"type:string;size:255;comment:镜像名称"`
	DockerImageId   string          `gorm:"type:string;size:255;comment:镜像ID"`
	SupportVersions SupportVersions `gorm:"type:text;not null;comment:可用版本"`
	IsInternal      bool            `gorm:"default:0;not null;comment:是否内置，云端内置驱动"`
	Language        string          `gorm:"type:string;size:255;comment:代码语言"`
	DeviceMarket
	OperateStatus string `gorm:"-"`
}

type DeviceMarket struct {
	Manual     string `gorm:"type:string;size:255;comment:驱动市场使用说明手册"`
	Icon       string `gorm:"type:text;comment:图标"`
	ClassifyId int    `gorm:"comment:分类"`
}

func (d DeviceLibrary) ConfigBody() string {
	return ""
}

func (d DeviceLibrary) GetConfig() (DeviceLibraryConfig, error) {
	body := d.ConfigBody()
	var dc DeviceLibraryConfig
	if err := json.Unmarshal([]byte(body), &dc); err != nil {
		return DeviceLibraryConfig{}, err
	}
	return dc, nil
}

func (d DeviceLibrary) GetConfigMap() (map[string]interface{}, error) {
	body := d.ConfigBody()
	var dc map[string]interface{}
	if err := json.Unmarshal([]byte(body), &dc); err != nil {
		return nil, err
	}
	return dc, nil
}

type SupportVersions []SupportVersion

type SupportVersion struct {
	Version            string
	ConfigFile         string
	IsDefault          bool // 是否为默认版本
	DockerParamsSwitch bool
	DockerParams       string
	ExpertMode         bool
	ExpertModeContent  string
	ConfigJson         string
}

func (c SupportVersions) Value() (driver.Value, error) {
	return GormValueWrap(c)
}

func (c *SupportVersions) Scan(value interface{}) error {
	return GormScanWrap(value, c)
}

func (table *DeviceLibrary) TableName() string {
	return "device_library"
}

func (table *DeviceLibrary) Get() interface{} {
	return *table
}

func (dl DeviceLibrary) DefaultVersion() SupportVersion {
	for _, v := range dl.SupportVersions {
		if v.Version == dl.Version {
			return v
		}
	}
	return SupportVersion{}
}

type DeviceLibraryConfig struct {
	DeviceServer map[string][]struct {
		Name         string      `json:"name"`
		Display      string      `json:"display"`
		Type         string      `json:"type"` // type 支持 int、string、float、bool、select、object、array
		DefaultValue string      `json:"defaultValue"`
		Options      interface{} `json:"options"`
	} `json:"deviceServer"`
	DeviceProtocols interface{}   `json:"deviceProtocols"`
	DeviceDpAttrs   []interface{} `json:"deviceDpAttrs"`
}

func GetLibrarySimpleBaseConfig() string {
	return `{"deviceServer": {},"deviceProtocols": {},"deviceDpAttrs": []}`
}
