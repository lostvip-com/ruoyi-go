package role

import (
	"lostvip.com/db"
	"time"
)

type SysRole struct {
	RoleId     int64     `json:"roleId" xorm:"not null pk autoincr comment('角色ID') BIGINT(20)"`
	RoleName   string    `json:"roleName" xorm:"not null comment('角色名称') VARCHAR(30)"`
	RoleKey    string    `json:"roleKey" xorm:"not null comment('角色权限字符串') VARCHAR(100)"`
	RoleSort   int       `json:"roleSort" xorm:"not null comment('显示顺序') INT(4)"`
	DataScope  string    `json:"dataScope" xorm:"default '1' comment('数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）') CHAR(1)"`
	Status     string    `json:"status" xorm:"not null comment('角色状态（0正常 1停用）') CHAR(1)"`
	DelFlag    string    `json:"delFlag" xorm:"default '0' comment('删除标志（0代表存在 2代表删除）') CHAR(1)"`
	CreateBy   string    `json:"createBy" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"createTime" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"updateBy" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"updateTime" xorm:"comment('更新时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
}

// 映射数据表
func (r *SysRole) TableName() string {
	return "sys_role"
}

// 插入数据
func (r *SysRole) Insert() (int64, error) {
	return db.GetInstance().Engine().Table("sys_role").Insert(r)
}

// 更新数据
func (r *SysRole) Update() (int64, error) {
	return db.GetInstance().Engine().Table("sys_role").ID(r.RoleId).Update(r)
}

// 删除
func (r *SysRole) Delete() (int64, error) {
	return db.GetInstance().Engine().Table("sys_role").ID(r.RoleId).Delete(r)
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *SysRole) FindOne() (bool, error) {
	return db.GetInstance().Engine().Table("sys_role").Get(r)
}
