package models

import (
	"github.com/lostvip-com/lv_framework/lv_db"
)

type SysDept struct {
	DeptId    int64  `gorm:"type:bigint(20);primary_key;auto_increment;部门id;" json:"deptId"`
	ParentId  int64  `gorm:"type:bigint(20);comment:父部门id;" json:"parentId"`
	Ancestors string `gorm:"type:varchar(50);comment:祖级列表;index:idx_ancestors" json:"ancestors"`
	DeptName  string `gorm:"type:varchar(30);comment:部门名称;" json:"deptName"`
	OrderNum  int    `gorm:"type:int(10);comment:显示顺序;" json:"orderNum"`
	Leader    string `gorm:"type:varchar(20);comment:负责人;" json:"leader"`
	Phone     string `gorm:"type:varchar(11);comment:联系电话;" json:"phone"`
	DeptType  string `gorm:"type:varchar(50);default:0;comment:组织类型;" json:"deptType"`
	Status    string `gorm:"type:char(1);comment:部门状态（0正常 1停用）;" json:"status"`
	UpdateBy  string `gorm:"type:varchar(64);comment:更新者;" json:"updateBy"`
	CreateBy  string `gorm:"type:varchar(32);comment:创建人;column:create_by;"  json:"createBy"`
	TenantId  int64  `gorm:"type:bigint(20);comment:租户id;" json:"tenantId" form:"tenantId"`
	BaseModel
	ParentName string `gorm:"-" json:"parentName"`
}

// 映射数据表
func (e *SysDept) TableName() string {
	return "sys_dept"
}

// 增
func (e *SysDept) Save() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

// 查
func (e *SysDept) FindById(id int64) (*SysDept, error) {
	err := lv_db.GetMasterGorm().Take(e, id).Error
	return e, err
}

// 查第一条
func (e *SysDept) FindOne() error {
	tb := lv_db.GetMasterGorm()
	if e.Ancestors != "" {
		tb = tb.Where("ancestors=?", e.Ancestors)
	}
	if e.DeptName != "" {
		tb = tb.Where("dept_name=?", e.DeptName)
	}
	if e.DeptId != 0 {
		tb = tb.Where("dept_id=?", e.DeptId)
	}
	err := tb.First(e).Error
	return err
}

// 改
func (e *SysDept) Update() error {
	return lv_db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}
func (e *SysDept) UpdateDelFlag(id int64) error {
	return lv_db.GetMasterGorm().Table(e.TableName()).Where("dept_id=?", id).Update("del_flag", "1").Error
}

// 删
func (e *SysDept) Delete() error {
	return lv_db.GetMasterGorm().Table(e.TableName()).Delete(e).Error
}
