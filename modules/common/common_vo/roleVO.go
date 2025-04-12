package common_vo

import (
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"time"
)

type SysRoleFlag struct {
	RoleId     int64     `json:"roleId" `
	RoleName   string    `json:"roleName" `
	RoleKey    string    `json:"roleKey" `
	RoleSort   int       `json:"roleSort" `
	DataScope  string    `json:"dataScope"`
	Status     string    `json:"status" `
	DelFlag    string    `json:"delFlag"`
	CreateBy   string    `json:"createBy" `
	CreateTime time.Time `json:"createTime" `
	UpdateBy   string    `json:"updateBy" `
	UpdateTime time.Time `json:"updateTime" `
	Remark     string    `json:"remark"`
	Flag       bool      `json:"flag" `
}

// DataScopeReq 数据权限保存请求参数
type DataScopeReq struct {
	RoleId    int64  `form:"roleId"  binding:"required"`
	RoleName  string `form:"roleName"  binding:"required"`
	RoleKey   string `form:"roleKey"  binding:"required"`
	DataScope string `form:"dataScope"  binding:"required"`
	DeptIds   string `form:"deptIds"`
}

// RolePageReq 分页请求参数
type RolePageReq struct {
	RoleName  string `form:"roleName"`  //角色名称
	Status    string `form:"status"`    //状态
	RoleKey   string `form:"roleKey"`   //角色键
	DataScope string `form:"dataScope"` //数据范围
	BeginTime string `form:"beginTime"` //开始时间
	EndTime   string `form:"endTime"`   //结束时间
	PageNum   int    `form:"pageNum"`   //当前页码
	PageSize  int    `form:"pageSize"`  //每页数
	SortName  string `form:"sortName"`  //排序字段
	SortOrder string `form:"sortOrder"` //排序方式
	//
	TenantId int64 `form:"tenantId"`
	lv_dto.Paging
}

// AddRoleReq 新增页面请求参数
type AddRoleReq struct {
	RoleName string `form:"roleName"  binding:"required"`
	RoleKey  string `form:"roleKey"  binding:"required"`
	RoleSort string `form:"roleSort"  binding:"required"`
	Status   string `form:"status"`
	Remark   string `form:"remark"`
	MenuIds  string `form:"menuIds""`
}

// EditRoleReq 修改页面请求参数
type EditRoleReq struct {
	RoleId   int64  `form:"roleId" binding:"required"`
	RoleName string `form:"roleName"  binding:"required"`
	RoleKey  string `form:"roleKey"  binding:"required"`
	RoleSort string `form:"roleSort"  binding:"required"`
	Status   string `form:"status"`
	Remark   string `form:"remark"`
	MenuIds  string `form:"menuIds"`
}
