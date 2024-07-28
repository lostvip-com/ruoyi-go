package cm_vo

import "github.com/lostvip-com/lv_framework/utils/lv_web"

// AddDictTypeReq 新增页面请求参数
type AddDictTypeReq struct {
	DictName string `form:"dictName"  binding:"required"`
	DictType string `form:"dictType"  binding:"required"`
	Status   string `form:"status"  binding:"required"`
	Remark   string `form:"remark"`
}

// EditDictTypeReq 修改页面请求参数
type EditDictTypeReq struct {
	DictId   int64  `form:"dictId" binding:"required"`
	DictName string `form:"dictName"  binding:"required"`
	DictType string `form:"dictType"  binding:"required"`
	Status   string `form:"status"  binding:"required"`
	Remark   string `form:"remark"`
}

// DictTypePageReq 分页请求参数
type DictTypePageReq struct {
	DictName  string `form:"dictName"`  //字典名称
	DictType  string `form:"dictType"`  //字典类型
	Status    string `form:"status"`    //字典状态
	BeginTime string `form:"beginTime"` //开始时间
	EndTime   string `form:"endTime"`   //结束时间
	lv_web.Paging
}

// 检查字典类型请求参数
type CheckDictTypeReq struct {
	DictId   int64  `form:"dictId"  binding:"required"`
	DictType string `form:"dictType"  binding:"required"`
}

// 检查字典类型请求参数
type CheckDictTypeALLReq struct {
	DictType string `form:"dictType"  binding:"required"`
}
