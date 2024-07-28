package model

import (
	"common/cm_model"
	"github.com/lostvip-com/lv_framework/db"
)

type SysRole struct {
	RoleId    int64  `gorm:"type:bigint(20);primary_key;auto_increment;角色ID;" json:"roleId"`
	RoleName  string `gorm:"type:varchar(30);comment:角色名称;" json:"roleName"`
	RoleKey   string `gorm:"type:varchar(100);comment:角色权限字符串;uniqueIndex:idx_roleKey" json:"roleKey"`
	RoleSort  int    `gorm:"type:int(11);comment:显示顺序;" json:"roleSort"`
	DataScope string `gorm:"type:char(1);comment:数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）;" json:"dataScope"`
	Status    string `gorm:"type:char(1);comment:角色状态（0正常 1停用）;" json:"status"`
	UpdateBy  string `gorm:"type:varchar(64);comment:更新者;" json:"updateBy"`
	CreateBy  string `gorm:"type:varchar(64);comment:创建者;" json:"updateBy"`
	Remark    string `gorm:"type:varchar(500);comment:备注;" json:"remark"`
	cm_model.BaseModel
}

// 映射数据表
func (r *SysRole) TableName() string {
	return "sys_role"
}

// 插入数据
func (r *SysRole) Insert() error {
	return db.GetMasterGorm().Save(r).Error
}

// 更新数据
func (r *SysRole) Update() error {
	return db.GetMasterGorm().Updates(r).Error
}

// 删除
func (r *SysRole) Delete() error {
	return db.GetMasterGorm().Delete(r).Error
}

// 根据结构体中已有的非空数据来获得单条数据
func (e *SysRole) FindOne() error {
	tb := db.GetMasterGorm()
	if e.RoleId != 0 {
		tb = tb.Where("role_id=? and del_flag=0", e.RoleId)
	}
	if e.RoleKey != "" {
		tb = tb.Where("role_key=? and del_flag=0", e.RoleKey)
	}
	err := tb.First(e).Error
	return err
}
