package cm_vo

import "github.com/lostvip-com/lv_framework/utils/lv_web"

// DeptPageReq 分页请求参数
type DeptPageReq struct {
	ParentId  int64  `form:"parentId"` //父部门ID
	NodeType  string `form:"nodeType"`
	DeptName  string `form:"deptName"`  //部门名称
	Status    string `form:"status"`    //状态
	BeginTime string `form:"beginTime"` //开始时间
	EndTime   string `form:"endTime"`   //结束时间
	TenantId  int64  `form:"tenantId"`  //结束时间
	lv_web.Paging
}

// AddDeptReq 新增页面请求参数
type AddDeptReq struct {
	ParentId int64  `form:"parentId"`
	NodeType string `form:"nodeType"`
	DeptName string `form:"deptName"  binding:"required"`
	OrderNum int    `form:"orderNum" `
	Leader   string `form:"leader"`
	Phone    string `form:"phone"`
	Status   string `form:"status"`
	TenantId int64  `form:"tenantId"` //结束时间
}

// EditDeptReq 修改页面请求参数
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
