package dept

import (
	"lostvip.com/db"
	"time"
)

type SysDept struct {
	DeptId     int64     `json:"deptId" xorm:"not null pk autoincr comment('部门id') BIGINT(20)"`
	ParentId   int64     `json:"parentId" xorm:"default 0 comment('父部门id') BIGINT(20)"`
	Ancestors  string    `json:"ancestors" xorm:"default '' comment('祖级列表') VARCHAR(50)"`
	DeptName   string    `json:"deptName" xorm:"default '' comment('部门名称') VARCHAR(30)"`
	OrderNum   int       `json:"orderNum" xorm:"default 0 comment('显示顺序') INT(4)"`
	Leader     string    `json:"leader" xorm:"comment('负责人') VARCHAR(20)"`
	Phone      string    `json:"phone" xorm:"comment('联系电话') VARCHAR(11)"`
	Email      string    `json:"email" xorm:"comment('邮箱') VARCHAR(50)"`
	Status     string    `json:"status" xorm:"default '0' comment('部门状态（0正常 1停用）') CHAR(1)"`
	DelFlag    string    `json:"delFlag" xorm:"default '0' comment('删除标志（0代表存在 1代表删除）') CHAR(1)"`
	CreateBy   string    `json:"createBy" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"createTime" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"updateBy" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"updateTime" xorm:"comment('更新时间') DATETIME"`

	TenantId int64 `json:"tenantId" xorm:"default 0 comment('租户id') BIGINT(20)"`
}

// 映射数据表
func TableName() string {
	return "sys_dept"
}

// 插入数据
func (r *SysDept) Insert() (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).Insert(r)
}

// 更新数据
func (r *SysDept) Update() (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).ID(r.DeptId).Update(r)
}

// 删除
func (r *SysDept) Delete() (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).ID(r.DeptId).Delete(r)
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *SysDept) FindOne() (bool, error) {
	return db.GetInstance().Engine().Table(TableName()).Get(r)
}
