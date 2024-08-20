// ==========================================================================
// LV自动生成model扩展代码列表 按需修改
// 生成日期：2024-07-19 17:16:13 +0800 CST
// 生成人：lv
// ==========================================================================
package vo

import (
	"github.com/lostvip-com/lv_framework/web/lv_dto"
)

// 新增页面请求参数
type AddIotProductReq struct {
	Name        string `form:"name"binding:"required" `
	Code        string `form:"code"`
	Protocol    string `form:"protocol"`
	NodeType    string `form:"nodeType"`
	Factory     string `form:"factory"`
	Description string `form:"description"`
	CreateBy    string
}

// 修改页面请求参数
type EditIotProductReq struct {
	Id          int64  `form:"id" binding:"required"`
	Name        string `form:"name"binding:"required" `
	Code        string `form:"code"`
	Protocol    string `form:"protocol"`
	NodeType    string `form:"nodeType"`
	NetType     string `form:"netType"`
	Factory     string `form:"factory"`
	Description string `form:"description"`
	UpdateBy    string
}

// 分页请求参数
type PageIotProductReq struct {
	Factory   string `form:"factory"`
	Name      string `form:"name"` //名字
	Code      string `form:"code"`
	BeginTime string `form:"beginTime"` //开始时间
	EndTime   string `form:"endTime"`   //结束时间
	lv_dto.Paging
}
