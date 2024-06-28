package model

import (
	"github.com/lv_framework/db"
	"time"
)

type SysRole struct {
	//RoleId     int64     `json:"roleId" xorm:"not null pk autoincr comment('角色ID') BIGINT(20)"`
	//RoleName   string    `json:"roleName" xorm:"not null comment('角色名称') VARCHAR(30)"`
	//RoleKey    string    `json:"roleKey" xorm:"not null comment('角色权限字符串') VARCHAR(100)"`
	//RoleSort   int       `json:"roleSort" xorm:"not null comment('显示顺序') INT(4)"`
	//DataScope  string    `json:"dataScope" xorm:"default '1' comment('数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）') CHAR(1)"`
	//Status     string    `json:"status" xorm:"not null comment('角色状态（0正常 1停用）') CHAR(1)"`
	//DelFlag    string    `json:"delFlag" xorm:"default '0' comment('删除标志（0代表存在 2代表删除）') CHAR(1)"`
	//CreateBy   string    `json:"createBy" xorm:"default '' comment('创建者') VARCHAR(64)"`
	//CreateTime time.Time `json:"createTime" xorm:"comment('创建时间') DATETIME"`
	//UpdateBy   string    `json:"updateBy" xorm:"default '' comment('更新者') VARCHAR(64)"`
	//UpdateTime time.Time `json:"updateTime" xorm:"comment('更新时间') DATETIME"`
	//Remark     string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
	RoleId     int64     `gorm:"type:bigint(20);primary_key;auto_increment;角色ID;" json:"roleId"`
	RoleName   string    `gorm:"type:varchar(30);comment:角色名称;" json:"roleName"`
	RoleKey    string    `gorm:"type:varchar(100);comment:角色权限字符串;uniqueIndex:idx_roleKey" json:"roleKey"`
	RoleSort   int       `gorm:"type:int(11);comment:显示顺序;" json:"roleSort"`
	DataScope  string    `gorm:"type:char(1);comment:数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）;" json:"dataScope"`
	Status     string    `gorm:"type:char(1);comment:角色状态（0正常 1停用）;" json:"status"`
	UpdateBy   string    `gorm:"type:varchar(64);comment:更新者;" json:"updateBy"`
	UpdateTime time.Time `gorm:"type:datetime;comment:更新时间;" json:"updateTime" time_format:"2006-01-02 15:04:05"`
	Remark     string    `gorm:"type:varchar(500);comment:备注;" json:"remark"`
	CreateTime time.Time `gorm:"type:datetime;comment:创建时间;column:create_time;" json:"createTime" time_format:"2006-01-02 15:04:05"`
	CreateBy   string    `gorm:"type:varchar(32);comment:创建人;column:create_by;"  json:"createBy"`
	DelFlag    int       `gorm:"type:tinyint(1);default:0;comment:删除标记;column:del_flag;" json:"delFlag"`
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
		tb = tb.Where("role_id=?", e.RoleId)
	}
	if e.RoleKey != "" {
		tb = tb.Where("role_key=?", e.RoleKey)
	}
	err := tb.First(e).Error
	return err
}
