package user_role

import (
	"lostvip.com/db"
	"xorm.io/core"
)

type SysUserRole struct {
	UserId int64 `json:"userId" xorm:"not null pk comment('用户ID') BIGINT(20)"`
	RoleId int64 `json:"roleId" xorm:"not null pk comment('角色ID') BIGINT(20)"`
}

// 映射数据表
func (r *SysUserRole) TableName() string {
	return "sys_user_role"
}

// 插入数据
func (r *SysUserRole) Insert() (int64, error) {
	return db.GetInstance().Engine().Table(r.TableName()).Insert(r)
}

// 更新数据
func (r *SysUserRole) Update() (int64, error) {
	return db.GetInstance().Engine().Table(r.TableName()).ID(core.PK{r.UserId, r.RoleId}).Update(r)
}

// 删除
func (r *SysUserRole) Delete() (int64, error) {
	return db.GetInstance().Engine().Table(r.TableName()).ID(core.PK{r.UserId, r.RoleId}).Delete(r)
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *SysUserRole) FindOne() (bool, error) {
	return db.GetInstance().Engine().Table(r.TableName()).Get(r)
}
