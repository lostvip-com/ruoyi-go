// ==========================================================================
// LV自动生成model扩展代码列表 按需修改
// 生成日期：2024-07-30 21:59:38 +0800 CST
// 生成人：lv
// ==========================================================================
package vo

import (
	"github.com/lostvip-com/lv_framework/web/lv_dto"
)

// 新增页面请求参数
type AddIotPrdPropertyReq struct {
	ProductId  int64  `form:"productId"`
	Name       string `form:"name"binding:"required" `
	Code       string `form:"code"`
	Tag        string `form:"tag"`
	AccessMode string `form:"accessMode"`
	Type       string `form:"type"`
	Unit       string `form:"unit"`
	Remark     string `form:"remark"`
	DelFlag    string `form:"delFlag"`
	TenantId   int64  `form:"tenantId"`
	CreateBy   string
}

// 修改页面请求参数
type EditIotPrdPropertyReq struct {
	Id         int64  `form:"id" binding:"required"`
	ProductId  int64  `form:"productId"`
	Name       string `form:"name"binding:"required" `
	Code       string `form:"code"`
	Tag        string `form:"tag"`
	AccessMode string `form:"accessMode"`
	Type       string `form:"type"`
	Unit       string `form:"unit"`
	Remark     string `form:"remark"`
	DelFlag    string `form:"delFlag"`
	TenantId   int64  `form:"tenantId"`
	UpdateBy   string
}

// 分页请求参数
type PageIotPrdPropertyReq struct {
	ProductId   int64  `form:"productId"`   //产品ID
	Name        string `form:"name"`        //名字
	Code        string `form:"code"`        //标识符
	Tag         string `form:"tag"`         //标签
	BeginTime   string `form:"beginTime"`   //开始时间
	EndTime     string `form:"endTime"`     //结束时间
	SearchValue string `form:"searchValue"` //结束时间
	lv_dto.Paging
}
