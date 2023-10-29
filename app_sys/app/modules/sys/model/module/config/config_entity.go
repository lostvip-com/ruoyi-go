// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2021-06-29 23:13:41 +0800 CST
// 生成路径: app/model/modules/config/SysConfig.go
// 生成人：robnote
// ==========================================================================

package config

import (
	"lostvip.com/db"
	"time"
)

// 数据表映射结构体
type SysConfig struct {
	ConfigId    int64     `json:"config_id" xorm:"not null pk autoincr comment('参数主键') int"`
	ConfigName  string    `json:"config_name" xorm:"comment('参数名称') varchar(100)"`
	ConfigKey   string    `json:"config_key" xorm:"comment('参数键名') varchar(100)"`
	ConfigValue string    `json:"config_value" xorm:"comment('参数键值') varchar(500)"`
	ConfigType  string    `json:"config_type" xorm:"comment('系统内置（Y是 N否）') char(1)"`
	CreateBy    string    `json:"create_by" xorm:"comment('创建者') varchar(64)"`
	CreateTime  time.Time `json:"create_time" xorm:"comment('创建时间') datetime"`
	UpdateBy    string    `json:"update_by" xorm:"comment('更新者') varchar(64)"`
	UpdateTime  time.Time `json:"update_time" xorm:"comment('更新时间') datetime"`
	Remark      string    `json:"remark" xorm:"comment('备注') varchar(500)"`
}

// 映射数据表
func TableName() string {
	return "sys_config"
}

// 插入数据
func (e *SysConfig) Insert() (int64, error) {
	return db.Instance().Engine().Table(TableName()).Insert(e)
}

// 更新数据
func (e *SysConfig) Update() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(e.ConfigId).Update(e)
}

// 删除
func (e *SysConfig) Delete() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(e.ConfigId).Delete(e)
}

// 批量删除
func DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(TableName()).In("config_id", ids).Delete(new(SysConfig))
}

// 根据结构体中已有的非空数据来获得单条数据
func (e *SysConfig) FindOne() (bool, error) {
	return db.Instance().Engine().Table(TableName()).Get(e)
}

// 根据条件查询
func Find(where, order string) ([]SysConfig, error) {
	var list []SysConfig
	err := db.Instance().Engine().Table(TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

// 指定字段集合查询
func FindIn(column string, args ...interface{}) ([]SysConfig, error) {
	var list []SysConfig
	err := db.Instance().Engine().Table(TableName()).In(column, args).Find(&list)
	return list, err
}

// 排除指定字段集合查询
func FindNotIn(column string, args ...interface{}) ([]SysConfig, error) {
	var list []SysConfig
	err := db.Instance().Engine().Table(TableName()).NotIn(column, args).Find(&list)
	return list, err
}
