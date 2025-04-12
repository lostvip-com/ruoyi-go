package model

import (
	"github.com/lostvip-com/lv_framework/lv_db"
)

type SysRoleDept struct {
	RoleId int64 `gorm:"type:bigint(20);primary_key;角色id;" json:"roleId"`
	DeptId int64 `gorm:"type:bigint(20);primary_key;部门id;" json:"deptId"`
}

// 映射数据表
func (r *SysRoleDept) TableName() string {
	return "sys_role_dept"
}

// 查
func (e *SysRoleDept) FindById() error {
	err := lv_db.GetMasterGorm().Take(e, "role_id=? and dept_id=?", e.RoleId, e.DeptId).Error
	return err
}

// 查第一条
func (e *SysRoleDept) FindOne() error {
	tb := lv_db.GetMasterGorm()
	if e.RoleId != 0 && e.DeptId != 0 {
		tb = tb.Where("role_id=? and dept_id=?", e.RoleId, e.DeptId)
	}
	err := tb.First(e).Error
	return err
}

// 改
func (e *SysRoleDept) Update() error {
	return lv_db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}

// 插入数据
func (r *SysRoleDept) Insert() error {
	return lv_db.GetMasterGorm().Save(r).Error
}

// 删除
func (r *SysRoleDept) Delete() error {
	return lv_db.GetMasterGorm().Delete(r).Error
}
