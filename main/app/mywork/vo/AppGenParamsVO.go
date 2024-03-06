// ==========================================================================
// LV自动生成model扩展代码列表 按需修改
// 生成日期：2024-02-28 14:21:50 +0800 CST
// 生成人：lv
// ==========================================================================
package vo

import (
	"lostvip.com/utils/lv_web"
)

// 新增页面请求参数
type AddAppGenParamsReq struct {
	ParamNo       int    `form:"paramNo"`
	ParamName     string `form:"paramName"binding:"required" `
	ParamType     string `form:"paramType"`
	Unit          string `form:"unit"`
	Remark        string `form:"remark"`
	MonitorTypeId int    `form:"monitorTypeId"`
	CreateBy      string
}

// 修改页面请求参数
type EditAppGenParamsReq struct {
	Id            int64  `form:"id" binding:"required"`
	ParamNo       int    `form:"paramNo"`
	ParamName     string `form:"paramName"binding:"required" `
	ParamType     string `form:"paramType"`
	Unit          string `form:"unit"`
	Remark        string `form:"remark"`
	MonitorTypeId int    `form:"monitorTypeId"`
	UpdateBy      string
}

// 分页请求参数
type PageAppGenParamsReq struct {
	MonitorTypeId int    `form:"monitorTypeId"` //参量号
	ParamNoStart  int    `form:"paramNoStart"`  //参量号
	ParamNoEnd    int    `form:"paramNoEnd"`    //参量号
	ParamName     string `form:"paramName"`     //参量名称
	ParamType     string `form:"paramType"`     //参量类型
	BeginTime     string `form:"beginTime"`     //开始时间
	EndTime       string `form:"endTime"`       //结束时间
	Remark        string `form:"remark"`        //排序
	SortOrder     string `form:"sortOrder"`     //排序
	UseFlag       string `form:"useFlag"`       //排序
	lv_web.Paging
}
