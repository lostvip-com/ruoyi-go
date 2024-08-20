package model

import (
	"github.com/lostvip-com/lv_framework/lv_db"
	"time"
)

type SysUserOnline struct {
	SessionId      string    `gorm:"type:varchar(64);primary_key;参数主键;" json:"sessionId"`
	LoginName      string    `gorm:"type:varchar(64);comment:登录名;" json:"loginName"`
	DeptName       string    `gorm:"type:varchar(64);comment:登录名;" json:"deptName"`
	Ipaddr         string    `gorm:"type:varchar(64);" json:"ipaddr"`
	LoginLocation  string    `gorm:"type:varchar(64);" json:"loginLocation"`
	Browser        string    `gorm:"type:varchar(64);" json:"browser"`
	Os             string    `gorm:"type:varchar(64);" json:"os"`
	Status         string    `gorm:"type:varchar(64);" json:"status"`
	StartTimestamp time.Time `gorm:"type:datetime;comment:开始时间;column:start_timestamp;" json:"startTimestamp" time_format:"2006-01-02 15:04:05"`
	LastAccessTime time.Time `gorm:"type:datetime;comment:上次访问;column:last_access_time;" json:"lastAccessTime" time_format:"2006-01-02 15:04:05"`
	ExpireTime     int       `gorm:"type:int;" json:"expireTime"`
	CreateTime     time.Time `gorm:"type:datetime;comment:创建时间;column:create_time;" json:"createTime" time_format:"2006-01-02 15:04:05"`
}

// 映射数据表
func (e *SysUserOnline) TableName() string {
	return "sys_user_online"
}

// 增
func (e *SysUserOnline) Save() error {
	return lv_db.GetMasterGorm().Create(e).Error
}

// 查
func (e *SysUserOnline) FindById() error {
	err := lv_db.GetMasterGorm().Take(e, e.SessionId).Error
	return err
}

// 改
func (e *SysUserOnline) Update() error {
	return lv_db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}

// 删
func (e *SysUserOnline) Delete() error {
	return lv_db.GetMasterGorm().Table(e.TableName()).Delete(e).Error
}
