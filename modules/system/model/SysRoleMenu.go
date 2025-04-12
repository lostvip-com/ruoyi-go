package model

import (
	"github.com/lostvip-com/lv_framework/lv_db"
)

type SysRoleMenu struct {
	RoleId int64 `gorm:"type:bigint(20);primary_key;comment:用户ID;"`
	MenuId int64 `gorm:"type:bigint(20);primary_key;comment:菜单ID;"`
}

// 映射数据表
func (r *SysRoleMenu) TableName() string {
	return "sys_role_menu"
}

// 插入数据
func (r *SysRoleMenu) Insert() error {
	return lv_db.GetMasterGorm().Save(r).Error
}

// 增
func (e *SysRoleMenu) Save() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

// 查
func (e *SysRoleMenu) FindById() error {
	err := lv_db.GetMasterGorm().Take(e, e.RoleId).Error
	return err
}

// 查第一条
func (e *SysRoleMenu) FindOne() error {
	tb := lv_db.GetMasterGorm()
	if e.MenuId != 0 && e.RoleId != 0 {
		tb = tb.Where("role_id=? and menu_id=?", e.RoleId, e.MenuId)
	}
	err := tb.First(e).Error
	return err
}

// 改
func (e *SysRoleMenu) Updates() error {
	return lv_db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}
