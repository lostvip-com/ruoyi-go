package vo

import (
	"lostvip.com/utils/lv_web"
	"time"
)

// Fill with you ideas below.

// 修改用户资料请求参数
type ProfileReq struct {
	UserName    string `form:"userName"  binding:"required,min=5,max=30"`
	Phonenumber string `form:"phonenumber"  binding:"required,len=10"`
	Email       string `form:"email"  binding:"required,email"`
	Sex         string `form:"sex"  binding:"required"`
}

// 修改密码请求参数
type PasswordReq struct {
	OldPassword string `form:"oldPassword" binding:"required,min=5,max=30"`
	NewPassword string `form:"newPassword" binding:"required,min=5,max=30"`
	Confirm     string `form:"confirm" binding:"required,min=5,max=30"`
}

// 重置密码请求参数
type ResetPwdReq struct {
	UserId   int64  `form:"userId"  binding:"required"`
	Password string `form:"password" binding:"required,min=5,max=30"`
}

// 检查email请求参数
type CheckEmailReq struct {
	UserId int64  `form:"userId"  binding:"required,min=5,max=30"`
	Email  string `form:"email"  binding:"required,email"`
}

// 检查email请求参数
type CheckEmailAllReq struct {
	Email string `form:"email"  binding:"required,email"`
}

// 检查phone请求参数
type CheckLoginNameReq struct {
	LoginName string `form:"loginName"  binding:"required"`
}

// 检查phone请求参数
type CheckPhoneReq struct {
	UserId      int64  `form:"userId"  binding:"required`
	Phonenumber string `form:"phonenumber"  binding:"required,len=11"`
}

// 检查phone请求参数
type CheckPhoneAllReq struct {
	Phonenumber string `form:"phonenumber"  binding:"required,len=11"`
}

// 检查密码请求参数
type CheckPasswordReq struct {
	Password string `form:"password"  binding:"required"`
}

// 查询用户列表请求参数
type SelectUserPageReq struct {
	LoginName   string `form:"loginName"`     //登录名
	Status      string `form:"status"`        //状态
	Phonenumber string `form:"phonenumber"`   //手机号码
	BeginTime   string `form:"beginTime"`     //数据范围
	EndTime     string `form:"endTime"`       //开始时间
	DeptId      int64  `form:"deptId"`        //结束时间
	SortName    string `form:"orderByColumn"` //排序字段
	SortOrder   string `form:"isAsc"`         //排序方式
	Ancestors   string `form:"ancestors"`     //排序方式
	//
	TenantId int64 `form:"tenantId"`
	lv_web.Paging
}

// 用户列表数据结构
type UserListEntity struct {
	UserId      int64     `json:"userId" xorm:"not null pk autoincr comment('用户ID') BIGINT(20)"`
	DeptId      int64     `json:"deptId" xorm:"comment('部门ID') BIGINT(20)"`
	LoginName   string    `json:"loginName" xorm:"not null comment('登录账号') VARCHAR(30)"`
	UserName    string    `json:"userName" xorm:"not null comment('用户昵称') VARCHAR(30)"`
	UserType    string    `json:"userType" xorm:"default '00' comment('用户类型（00系统用户）') VARCHAR(2)"`
	Email       string    `json:"email" xorm:"default '' comment('用户邮箱') VARCHAR(50)"`
	Phonenumber string    `json:"phonenumber" xorm:"default '' comment('手机号码') VARCHAR(11)"`
	Sex         string    `json:"sex" xorm:"default '0' comment('用户性别（0男 1女 2未知）') CHAR(1)"`
	Avatar      string    `json:"avatar" xorm:"default '' comment('头像路径') VARCHAR(100)"`
	Password    string    `json:"password" xorm:"default '' comment('密码') VARCHAR(50)"`
	Salt        string    `json:"salt" xorm:"default '' comment('盐加密') VARCHAR(20)"`
	Status      string    `json:"status" xorm:"default '0' comment('帐号状态（0正常 1停用）') CHAR(1)"`
	DelFlag     string    `json:"delFlag" xorm:"default '0' comment('删除标志（0代表存在 2代表删除）') CHAR(1)"`
	LoginIp     string    `json:"loginIp" xorm:"default '' comment('最后登录IP') VARCHAR(50)"`
	LoginDate   time.Time `json:"loginDate" xorm:"comment('最后登录时间') DATETIME"`
	CreateBy    string    `json:"createBy" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime  time.Time `json:"createTime" xorm:"comment('创建时间') DATETIME"`
	UpdateBy    string    `json:"updateBy" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime  time.Time `json:"updateTime" xorm:"comment('更新时间') DATETIME"`
	Remark      string    `json:"remark" xorm:"comment('备注') VARCHAR(500)"`
	DeptName    string    `json:"deptName"` // 部门名称
	Leader      string    `json:"leader"`   // 负责人
}

// 新增用户资料请求参数
type AddReq struct {
	UserName    string `form:"userName"  binding:"required,min=5,max=30"`
	Phonenumber string `form:"phonenumber"  binding:"required,len=11"`
	Email       string `form:"email"  binding:"required,email"`
	LoginName   string `form:"loginName"  binding:"required"`
	Password    string `form:"password"  binding:"required,min=5,max=30"`
	DeptId      int64  `form:"deptId" binding:"required`
	Sex         string `form:"sex"  binding:"required"`
	Status      string `form:"status"`
	RoleIds     string `form:"roleIds"`
	PostIds     string `form:"postIds"`
	Remark      string `form:"remark"`
}

// 新增用户资料请求参数
type EditReq struct {
	UserId      int64  `form:"userId" binding:"required`
	UserName    string `form:"userName"  binding:"required,max=30"`
	Phonenumber string `form:"phonenumber"  binding:"required,len=11"`
	Email       string `form:"email"  binding:"required,email"`
	DeptId      int64  `form:"deptId" binding:"required`
	Sex         string `form:"sex"  binding:"required"`
	Status      string `form:"status"`
	RoleIds     string `form:"roleIds"`
	PostIds     string `form:"postIds"`
	Remark      string `form:"remark"`
}
