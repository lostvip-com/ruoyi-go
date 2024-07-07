package model

import (
	"github.com/lv_framework/db"
)

// 用户和角色关联
type SysUserRole struct {
	UserId int64 `gorm:"type:bigint(20);primary_key;auto_increment;用户ID;" json:"userId"`
	RoleId int64 `gorm:"type:bigint(20);primary_key;auto_increment;角色ID;" json:"roleId"`
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
	tb := db.GetMasterGorm().Table(e.TableName())
	if e.UserId != 0 && e.RoleId != 0 {
		tb = tb.Table(e.TableName()).Where("role_id=? and user_id=?", e.RoleId, e.UserId)
	}
	err := tb.First(e).Error
	return err
}

// 删
func (e *SysUserRole) Delete() error {
	err := e.FindOne()
	if e == nil {
		return db.GetMasterGorm().Delete(e).Error
	} else {
		return err
	}

}
