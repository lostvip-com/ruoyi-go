package common_vo

import (
	"github.com/lostvip-com/lv_framework/web/lv_dto"
)

// Fill with you ideas below.

// 修改用户资料请求参数
type ProfileReq struct {
	UserName    string `form:"userName"  binding:"required,min=1,max=30"`
	Phonenumber string `form:"phonenumber"  binding:"required,len=11"`
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
	LoginName   string `form:"loginName"`   //登录名
	Status      string `form:"status"`      //状态
	Phonenumber string `form:"phonenumber"` //手机号码
	BeginTime   string `form:"beginTime"`   //数据范围
	EndTime     string `form:"endTime"`     //开始时间
	DeptId      int64  `form:"deptId"`      //结束时间
	SortName    string `form:"sortName"`    //排序字段
	SortOrder   string `form:"sortOrder"`   //排序方式
	Ancestors   string `form:"ancestors"`   //排序方式
	//
	TenantId int64 `form:"tenantId"`
	lv_dto.Paging
}

// 新增用户资料请求参数
type AddUserReq struct {
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
type EditUserReq struct {
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
