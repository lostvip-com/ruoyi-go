package model

import (
	"errors"
	"github.com/lostvip-com/lv_framework/lv_db"
	"time"
)

// SysConfig 参数配置
type SysConfig struct {
	ConfigId    int64     `gorm:"type:int(11);primary_key;auto_increment;参数主键;" json:"configId"`
	ConfigName  string    `gorm:"type:varchar(100);comment:参数名称;" json:"configName"`
	ConfigKey   string    `gorm:"type:varchar(100);comment:参数键名;" json:"configKey"`
	ConfigValue string    `gorm:"type:varchar(500);comment:参数键值;" json:"configValue"`
	ConfigType  string    `gorm:"type:char(1);comment:系统内置（Y是 N否）;" json:"configType"`
	UpdateBy    string    `gorm:"type:varchar(64);comment:更新者;" json:"updateBy"`
	UpdateTime  time.Time `gorm:"type:datetime;comment:更新时间;" json:"updateTime" time_format:"2006-01-02 15:04:05"`
	Remark      string    `gorm:"type:varchar(500);comment:备注;" json:"remark"`
	CreateTime  time.Time `gorm:"type:datetime;comment:创建时间;column:create_time;" json:"createTime" time_format:"2006-01-02 15:04:05"`
	CreateBy    string    `gorm:"type:varchar(32);comment:创建人;column:create_by;"  json:"createBy"`
	DelFlag     int       `gorm:"type:tinyint(1);default:0;comment:删除标记;column:del_flag;" json:"delFlag"`
}

func (e *SysConfig) TableName() string {
	return "sys_config"
}

// 增
func (e *SysConfig) Save() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

// 查
func (e *SysConfig) FindById(id int64) (*SysConfig, error) {
	err := lv_db.GetMasterGorm().First(e, id).Error
	return e, err
}

// 查第一条
func (e *SysConfig) FindOne() error {
	tb := lv_db.GetMasterGorm().Table(e.TableName())
	if e.ConfigId != 0 {
		tb = tb.Where("config_id=?", e.ConfigId)
	}
	if e.ConfigKey != "" {
		tb = tb.Where("config_key=?", e.ConfigKey)
	}

	err := tb.First(e).Error
	if err != nil {
		return errors.New("未找到系统配置信息Key:" + e.ConfigKey)
	}
	return err
}

// 改
func (e *SysConfig) Update() error {
	return lv_db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}

// 删
func (e *SysConfig) Delete() error {
	return lv_db.GetMasterGorm().Table(e.TableName()).Delete(e).Error
}
