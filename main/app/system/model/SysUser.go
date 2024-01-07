// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期：2024-01-06 17:33:26 +0800 CST
// 生成人：lv
// ==========================================================================

package model

import (
	"lostvip.com/db"
	"time"
)

// SysUser 用户信息
type SysUser struct {
	UserId      int64     `gorm:"type:bigint(20);primary_key;auto_increment;用户ID;"     json:"userId"  form:"userId"`
	DeptId      int64     `gorm:"type:bigint(20);comment:部门ID;" json:"deptId" form:"deptId"`
	LoginName   string    `gorm:"type:varchar(30);comment:登录账号;" json:"loginName" form:"loginName"`
	UserName    string    `gorm:"type:varchar(30);comment:用户昵称;" json:"userName" form:"userName"`
	UserType    string    `gorm:"type:varchar(2);comment:用户类型（00系统用户）;" json:"userType" form:"userType"`
	Email       string    `gorm:"type:varchar(50);comment:用户邮箱;" json:"email" form:"email"`
	Phonenumber string    `gorm:"type:varchar(11);comment:手机号码;" json:"phonenumber" form:"phonenumber"`
	Sex         string    `gorm:"type:char(1);comment:用户性别（0男 1女 2未知）;" json:"sex" form:"sex"`
	Avatar      string    `gorm:"type:varchar(100);comment:头像路径;" json:"avatar" form:"avatar"`
	Password    string    `gorm:"type:varchar(50);comment:密码;" json:"password" form:"password"`
	Salt        string    `gorm:"type:varchar(20);comment:盐加密;" json:"salt" form:"salt"`
	Status      string    `gorm:"type:char(1);comment:帐号状态（0正常 1停用）;" json:"status" form:"status"`
	LoginIp     string    `gorm:"type:varchar(50);comment:最后登陆IP;" json:"loginIp" form:"loginIp"`
	LoginDate   time.Time `gorm:"type:datetime;comment:最后登陆时间;" json:"loginDate" form:"loginDate" time_format:"2006-01-02 15:04:05"`
	UpdateBy    string    `gorm:"type:varchar(64);comment:更新者;" json:"updateBy" form:"updateBy"`
	UpdateTime  time.Time `gorm:"type:datetime;comment:更新时间;" json:"updateTime" form:"updateTime" time_format:"2006-01-02 15:04:05"`
	Remark      string    `gorm:"type:varchar(500);comment:备注;" json:"remark" form:"remark"`
	TenantId    int64     `gorm:"type:bigint(20);comment:租户id;" json:"tenantId" form:"tenantId"`
	CreateTime  time.Time `gorm:"type:datetime;comment:创建时间;"  column:create_time; time_format:"2006-01-02 15:04:05" json:"createTime"  `
	CreateBy    string    `gorm:"type:varchar(32);comment:创建人;" column:create_by; json:"createBy" form:"createBy"`
	DelFlag     int       `gorm:"type:tinyint(1);default:0;comment:删除标记;" column:del_flag; json:"delFlag"`
}

// 映射数据表
func (e *SysUser) TableName() string {
	return "sys_user"
}

// 增
func (e *SysUser) Insert() error {
	return db.GetMasterGorm().Save(e).Error
}

// 查
func (e *SysUser) GetById() error {
	err := db.GetMasterGorm().Take(e).Error
	return err
}

// 查
func (e *SysUser) FindOne() error {
	err := db.GetMasterGorm().First(e).Error
	return err
}

// 改
func (e *SysUser) Updates() error {
	return db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}

// 删
func (e *SysUser) Delete() error {
	return db.GetMasterGorm().Table(e.TableName()).Delete(e).Error
}
