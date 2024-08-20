// ==========================================================================
// LV自动生成model扩展代码列表 按需修改
// 生成日期：2024-07-30 21:59:36 +0800 CST
// 生成人：lv
// ==========================================================================
package vo

import (
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"time"
)

// 新增页面请求参数
type AddIotPrdEventReq struct {
	ProductId    string `form:"productId"`
	EventType    string `form:"eventType"`
	Code         string `form:"code"`
	Tag          string `form:"tag"`
	Name         string `form:"name"binding:"required" `
	Remark       string `form:"remark"`
	OutputParams string `form:"outputParams"`
	DelFlag      string `form:"delFlag"`
	TenantId     int64  `form:"tenantId"`
	CreateBy     string
}

// 修改页面请求参数
type EditIotPrdEventReq struct {
	Id           int64  `form:"id" binding:"required"`
	ProductId    string `form:"productId"`
	EventType    string `form:"eventType"`
	Code         string `form:"code"`
	Tag          string `form:"tag"`
	Name         string `form:"name"binding:"required" `
	Remark       string `form:"remark"`
	OutputParams string `form:"outputParams"`
	DelFlag      string `form:"delFlag"`
	TenantId     int64  `form:"tenantId"`
	UpdateBy     string
}

// 分页请求参数
type PageIotPrdEventReq struct {
	ProductId int64  `form:"productId"` //产品ID
	EventType string `form:"eventType"` //事件类型
	Code      string `form:"code"`      //标识符
	Tag       string `form:"tag"`       //标签
	BeginTime string `form:"beginTime"` //开始时间
	EndTime   string `form:"endTime"`   //结束时间
	lv_dto.Paging
}

// 分页请求结果
type IotPrdEventResp struct {
	Id           int64     `json:"id"`
	ProductId    int64     `json:"productId"`
	EventType    string    `json:"eventType"`
	Code         string    `json:"code"`
	Tag          string    `json:"tag"`
	Name         string    `json:"name"`
	Remark       string    `json:"remark"`
	OutputParams string    `json:"outputParams"`
	DelFlag      string    `json:"delFlag"`
	UpdateBy     string    `json:"updateBy"`
	UpdateTime   time.Time `json:"updateTime" time_format:"2006-01-02 15:04:05"`
	CreateTime   time.Time `json:"createTime" time_format:"2006-01-02 15:04:05"`
	CreateBy     string    `json:"createBy"`
}
