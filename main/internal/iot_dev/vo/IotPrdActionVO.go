// ==========================================================================
// LV自动生成model扩展代码列表 按需修改
// 生成日期：2024-07-30 21:59:34 +0800 CST
// 生成人：lv
// ==========================================================================
package vo

import (
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"time"
)

// 新增页面请求参数
type AddIotPrdActionReq struct {
	ProductId    int    `form:"productId"`
	Code         string `form:"code"`
	Tag          string `form:"tag"`
	Name         string `form:"name"binding:"required" `
	CallType     string `form:"callType"`
	InputParams  string `form:"inputParams"`
	OutputParams string `form:"outputParams"`
	Remark       string `form:"remark"`
	DelFlag      string `form:"delFlag"`
	TenantId     int64  `form:"tenantId"`
	CreateBy     string
}

// 修改页面请求参数
type EditIotPrdActionReq struct {
	Id           int64  `form:"id" binding:"required"`
	ProductId    int64  `form:"productId"`
	Code         string `form:"code"`
	Tag          string `form:"tag"`
	Name         string `form:"name"binding:"required" `
	CallType     string `form:"callType"`
	InputParams  string `form:"inputParams"`
	OutputParams string `form:"outputParams"`
	Remark       string `form:"remark"`
	DelFlag      string `form:"delFlag"`
	TenantId     int64  `form:"tenantId"`
	UpdateBy     string
}

// 分页请求参数
type PageIotPrdActionReq struct {
	ProductId int    `form:"productId"` //产品ID
	Code      string `form:"code"`      //标识符
	Tag       string `form:"tag"`       //标签
	Name      string `form:"name"`      //名字
	BeginTime string `form:"beginTime"` //开始时间
	EndTime   string `form:"endTime"`   //结束时间
	lv_dto.Paging
}

// 分页请求结果
type IotPrdActionResp struct {
	Id           int       `json:"id"`
	ProductId    int       `json:"productId"`
	Code         string    `json:"code"`
	Tag          string    `json:"tag"`
	Name         string    `json:"name"`
	CallType     string    `json:"callType"`
	InputParams  string    `json:"inputParams"`
	OutputParams string    `json:"outputParams"`
	Remark       string    `json:"remark"`
	DelFlag      string    `json:"delFlag"`
	TenantId     int64     `json:"tenantId"`
	UpdateBy     string    `json:"updateBy"`
	UpdateTime   time.Time `json:"updateTime" time_format:"2006-01-02 15:04:05"`
	CreateTime   time.Time `json:"createTime" time_format:"2006-01-02 15:04:05"`
	CreateBy     string    `json:"createBy"`
}
