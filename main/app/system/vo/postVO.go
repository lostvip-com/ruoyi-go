package vo

import "lostvip.com/utils/lv_web"

// 新增页面请求参数
type AddPostReq struct {
	PostName string `form:"postName"  binding:"required"`
	PostCode string `form:"postCode"  binding:"required"`
	PostSort int    `form:"postSort"  binding:"required"`
	Status   string `form:"status"    binding:"required"`
	Remark   string `form:"remark"`
}

// 修改页面请求参数
type EditSysPostReq struct {
	PostId   int64  `form:"postId" binding:"required"`
	PostName string `form:"postName"  binding:"required"`
	PostCode string `form:"postCode"  binding:"required"`
	PostSort int    `form:"postSort"  binding:"required"`
	Status   string `form:"status"    binding:"required"`
	Remark   string `form:"remark"`
}

// 分页请求参数
type SelectPostPageReq struct {
	PostCode  string `form:"postCode"`  //岗位编码
	Status    string `form:"status"`    //状态
	PostName  string `form:"postName"`  //岗位名称
	BeginTime string `form:"beginTime"` //开始时间
	EndTime   string `form:"endTime"`   //结束时间
	SortName  string `form:"sortName"`  //排序字段
	SortOrder string `form:"sortOrder"` //排序方式
	Remark    string `form:"remark"`    //每页数
	lv_web.Paging
}

// 检查编码请求参数
type CheckPostCodeReq struct {
	PostId   int64  `form:"postId"  binding:"required"`
	PostCode string `form:"postCode"  binding:"required"`
}

// 检查编码请求参数
type CheckPostCodeALLReq struct {
	PostCode string `form:"postCode"  binding:"required"`
}

// 检查名称请求参数
type CheckPostNameReq struct {
	PostId   int64  `form:"postId"  binding:"required"`
	PostName string `form:"postName"  binding:"required"`
}

// 检查名称请求参数
type CheckPostNameALLReq struct {
	PostName string `form:"postName"  binding:"required"`
}
