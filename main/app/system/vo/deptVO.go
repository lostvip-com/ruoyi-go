package vo

import (
	"main/app/system/model"
)

// Fill with you ideas below.

// GenTable is the golang structure for table sys_dept.
type SysDeptExtend struct {
	model.SysDept `xorm:"extends"`
	ParentName    string `json:"parentName"`
}

// 分页请求参数
type DeptPageReq struct {
	ParentId  int64  `form:"parentId"`  //父部门ID
	DeptName  string `form:"deptName"`  //部门名称
	Status    string `form:"status"`    //状态
	PageNum   int    `form:"pageNum"`   //当前页码
	PageSize  int    `form:"pageSize"`  //每页数
	SortName  string `form:"sortName"`  //排序字段
	SortOrder string `form:"sortOrder"` //排序方式
	BeginTime string `form:"beginTime"` //开始时间
	EndTime   string `form:"endTime"`   //结束时间
	TenantId  int64  `form:"tenantId"`  //结束时间
}

// 新增页面请求参数
type AddDeptReq struct {
	ParentId int64  `form:"parentId"`
	DeptName string `form:"deptName"  binding:"required"`
	OrderNum int    `form:"orderNum" binding:"required"`
	Leader   string `form:"leader"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
	Status   string `form:"status"`
	TenantId int64  `form:"tenantId"` //结束时间
}

// 修改页面请求参数
type EditDeptReq struct {
	DeptId int64 `form:"deptId" binding:"required"`
	AddDeptReq
}

// 检查菜单名称请求参数
type CheckDeptNameReq struct {
	DeptId   int64  `form:"deptId"  binding:"required"`
	ParentId int64  `form:"parentId"  binding:"required"`
	DeptName string `form:"deptName"  binding:"required"`
}

// 检查菜单名称请求参数
type CheckDeptNameALLReq struct {
	ParentId int64  `form:"parentId"  binding:"required"`
	DeptName string `form:"deptName"  binding:"required"`
}
