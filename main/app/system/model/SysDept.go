package model

import (
	"github.com/lv_framework/db"
	"time"
)

type SysDept struct {
	DeptId     int64     `gorm:"type:bigint(20);primary_key;auto_increment;部门id;" json:"deptId"`
	ParentId   int64     `gorm:"type:bigint(20);comment:父部门id;" json:"parentId"`
	Ancestors  string    `gorm:"type:varchar(50);comment:祖级列表;uniqueIndex:idx_ancestors" json:"ancestors"`
	DeptName   string    `gorm:"type:varchar(30);comment:部门名称;" json:"deptName"`
	OrderNum   int       `gorm:"type:int(10);comment:显示顺序;" json:"orderNum"`
	Leader     string    `gorm:"type:varchar(20);comment:负责人;" json:"leader"`
	Phone      string    `gorm:"type:varchar(11);comment:联系电话;" json:"phone"`
	Email      string    `gorm:"type:varchar(50);comment:邮箱;" json:"email"`
	Status     string    `gorm:"type:char(1);comment:部门状态（0正常 1停用）;" json:"status"`
	UpdateBy   string    `gorm:"type:varchar(64);comment:更新者;" json:"updateBy"`
	UpdateTime time.Time `gorm:"type:datetime;comment:更新时间;" json:"updateTime" time_format:"2006-01-02 15:04:05"`
	TenantId   int64     `gorm:"type:bigint(20);comment:租户id;" json:"tenantId"`
	CreateTime time.Time `gorm:"type:datetime;comment:创建时间;column:create_time;" json:"createTime" time_format:"2006-01-02 15:04:05"`
	CreateBy   string    `gorm:"type:varchar(32);comment:创建人;column:create_by;"  json:"createBy"`
	DelFlag    int       `gorm:"type:tinyint(1);default:0;comment:删除标记;column:del_flag;" json:"delFlag"`
}

// 映射数据表
func (e *SysDept) TableName() string {
	return "sys_dept"
}

// 增
func (e *SysDept) Save() error {
	return db.GetMasterGorm().Save(e).Error
}

// 查
func (e *SysDept) FindById() error {
	err := db.GetMasterGorm().Take(e, e.DeptId).Error
	return err
}

// 查第一条
func (e *SysDept) FindOne() error {
	tb := db.GetMasterGorm()
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
	return db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}

// 删
func (e *SysDept) Delete() error {
	return db.GetMasterGorm().Table(e.TableName()).Delete(e).Error
}
