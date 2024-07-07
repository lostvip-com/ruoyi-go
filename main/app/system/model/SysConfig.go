package model

import (
	"github.com/lv_framework/db"
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
	return db.GetMasterGorm().Save(e).Error
}

// 查
func (e *SysConfig) FindById() error {
	err := db.GetMasterGorm().Take(e, e.ConfigId).Error
	return err
}

// 查第一条
func (e *SysConfig) FindOne() error {
	tb := db.GetMasterGorm().Table(e.TableName())
	if e.ConfigId != 0 {
		tb = tb.Where("config_id=?", e.ConfigId)
	}
	if e.ConfigKey != "" {
		tb = tb.Where("config_key=?", e.ConfigKey)
	}

	err := tb.First(e).Error
	return err
}

// 改
func (e *SysConfig) Update() error {
	return db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}

// 删
func (e *SysConfig) Delete() error {
	return db.GetMasterGorm().Table(e.TableName()).Delete(e).Error
}

//
//type SysConfig struct {
//	ConfigId    int64     `json:"configId" xorm:"not null pk autoincr comment('参数主键') INT(5)"`
//	ConfigName  string    `json:"configName" xorm:"default '' comment('参数名称') VARCHAR(100)"`
//	ConfigKey   string    `json:"configKey" xorm:"default '' comment('参数键名') VARCHAR(100)"`
//	ConfigValue string    `json:"configValue" xorm:"default '' comment('参数键值') VARCHAR(500)"`
//	ConfigType  string    `json:"configType" xorm:"default 'N' comment('系统内置（Y是 N否）') CHAR(1)"`
//	CreateBy    string    `json:"createBy" xorm:"default '' comment('创建者') VARCHAR(64)"`
//	CreateTime  time.Time `json:"createTime" xorm:"comment('创建时间') DATETIME"`
//	UpdateBy    string    `json:"updateBy" xorm:"default '' comment('更新者') VARCHAR(64)"`
//	UpdateTime  time.Time `json:"updateTime" xorm:"comment('更新时间') DATETIME"`
//	Remark      string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
//}
//
//// 映射数据表
//func (r *SysConfig) TableName() string {
//	return "sys_config"
//}
//
//// 插入数据
//func (r *SysConfig) Insert() (int64, error) {
//	return db.GetInstance().Engine().Insert(r)
//}
//
//// 更新数据
//func (r *SysConfig) Update() (int64, error) {
//	return db.GetInstance().Engine().Table("sys_config").ID(r.ConfigId).Update(r)
//}
//
//// 删除
//func (r *SysConfig) Delete() (int64, error) {
//	return db.GetInstance().Engine().Table("sys_config").ID(r.ConfigId).Delete(r)
//}
//
//// 根据结构体中已有的非空数据来获得单条数据
//func (r *SysConfig) FindOne() (bool, error) {
//	return db.GetInstance().Engine().Table("sys_config").Get(r)
//}
//
//// 根据条件查询
//func (r *SysConfig) Find(where, order string) ([]SysConfig, error) {
//	var list []SysConfig
//	err := db.GetInstance().Engine().Table("sys_config").Where(where).OrderBy(order).Find(&list)
//	return list, err
//}
