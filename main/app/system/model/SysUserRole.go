package model

import (
	"lostvip.com/db"
	"time"
)

// SysUserRole 用户和角色关联
type SysUserRole struct {
	UserId     int64     `gorm:"type:bigint(20);primary_key;auto_increment;用户ID;" json:"userId"`
	RoleId     int64     `gorm:"type:bigint(20);primary_key;auto_increment;角色ID;" json:"roleId"`
	CreateTime time.Time `gorm:"type:datetime;comment:创建时间;column:create_time;" json:"createTime" time_format:"2006-01-02 15:04:05"`
	CreateBy   string    `gorm:"type:varchar(32);comment:创建人;column:create_by;"  json:"createBy"`
	DelFlag    int       `gorm:"type:tinyint(1);default:0;comment:删除标记;column:del_flag;" json:"delFlag"`
}

func (e *SysUserRole) TableName() string {
	return "sys_user_role"
}

// 增
func (e *SysUserRole) Save() error {
	return db.GetMasterGorm().Save(e).Error
}

// 查
func (e *SysUserRole) FindById() error {
	err := db.GetMasterGorm().Take(e, " user_id=? and role_id=? ", e.UserId, e.RoleId).Error
	return err
}

// 查第一条
func (e *SysUserRole) FindOne() error {
	tb := db.GetMasterGorm()
	if e.UserId != 0 && e.RoleId != 0 {
		tb = tb.Where("role_id=? and user_id=?", e.RoleId, e.UserId)
	}
	err := tb.First(e).Error
	return err
}

// 改
func (e *SysUserRole) Updates() error {
	return db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}

// 删
func (e *SysUserRole) Delete() error {
	return db.GetMasterGorm().Table(e.TableName()).Delete(e).Error
}
