package online

import (
	"github.com/lv_framework/db"
	"time"
)

type UserOnline struct {
	SessionId      string    `gorm:"type:varchar(64);primary_key;参数主键;" json:"sessionId"`
	LoginName      string    `gorm:"type:varchar(64);comment:登录名;" json:"login_name"`
	DeptName       string    `gorm:"type:varchar(64);comment:登录名;" json:"dept_name"`
	Ipaddr         string    `gorm:"type:varchar(64);" json:"ipaddr"`
	LoginLocation  string    `gorm:"type:varchar(64);" json:"login_location"`
	Browser        string    `gorm:"type:varchar(64);" json:"browser"`
	Os             string    `gorm:"type:varchar(64);" json:"os"`
	Status         string    `gorm:"type:varchar(64);" json:"status"`
	StartTimestamp time.Time `gorm:"type:datetime;comment:开始时间;column:start_timestamp;" json:"start_timestamp" time_format:"2006-01-02 15:04:05"`
	LastAccessTime time.Time `gorm:"type:datetime;comment:上次访问;column:last_access_time;" json:"last_access_time" time_format:"2006-01-02 15:04:05"`
	ExpireTime     int       `gorm:"type:int;" json:"expire_time"`
	CreateTime     time.Time `gorm:"type:datetime;comment:创建时间;column:create_time;" json:"createTime" time_format:"2006-01-02 15:04:05"`
}

// 映射数据表
func (e *UserOnline) TableName() string {
	return "sys_user_online"
}

// 增
func (e *UserOnline) Save() error {
	return db.GetMasterGorm().Save(e).Error
}

// 查
func (e *UserOnline) FindById() error {
	err := db.GetMasterGorm().Take(e, e.SessionId).Error
	return err
}

// 改
func (e *UserOnline) Update() error {
	return db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}

// 删
func (e *UserOnline) Delete() error {
	return db.GetMasterGorm().Table(e.TableName()).Delete(e).Error
}
