// ==========================================================================
// LV自动生成model扩展代码列表 按需修改
// 生成日期：2024-07-19 17:09:35 +0800 CST
// 生成人：lv
// ==========================================================================
package vo

import (
	"github.com/lostvip-com/lv_framework/web/lv_dto"
)

// 新增页面请求参数
type AddIotDeviceReq struct {
	ProductId       int64  `form:"productId"`
	DeptId          int64  `form:"deptId"`
	NodeType        string `form:"nodeType"`
	GatewayId       int64  `form:"gatewayId"` //网关ID
	Name            string `form:"name"binding:"required" `
	Sn              string `form:"sn"`
	Status          string `form:"status"`
	DevNo           string `form:"devNo"`
	Description     string `form:"description"`
	Secret          string `form:"secret"`
	InstallLocation string `form:"installLocation"`
	CreateBy        string
}

// 修改页面请求参数
type EditIotDeviceReq struct {
	Id              int64  `form:"id" binding:"required"`
	ProductId       int64  `form:"productId"`
	DeptId          int64  `form:"deptId"`
	NodeType        string `form:"nodeType"`
	GatewayId       int64  `form:"gatewayId"` //网关ID
	DriveInstanceId string `form:"driveInstanceId"`
	Name            string `form:"name"binding:"required" `
	Status          string `form:"status"`
	Sn              string `form:"sn"`
	DevNo           string `form:"devNo"`
	Description     string `form:"description"`
	Secret          string `form:"secret"`
	InstallLocation string `form:"installLocation"`
	UpdateBy        string
}

// 分页请求参数
type PageIotDeviceReq struct {
	//
	NodeTypeArr []string
	DeptId      int64  `form:"deptId"`
	ProductId   int64  `form:"productId"` //产品ID
	NodeType    string `form:"nodeType"`
	GatewayId   int64  `form:"gatewayId"` //网关ID
	Name        string `form:"name"`
	Sn          string `form:"sn"`
	Status      string `form:"status"`
	DevNo       string `form:"devNo"`
	BeginTime   string `form:"beginTime"` //开始时间
	EndTime     string `form:"endTime"`   //结束时间
	Ancestors   string `form:"ancestors"`
	GwExist     string `form:"gwExist"`
	lv_dto.Paging
}

type IotDeviceResp struct {
	Id              int64  `json:"id"`
	ProductId       int64  `json:"productId"`
	DeptId          int64  `json:"deptId"`
	DeptName        string `json:"deptName"`
	NodeType        string `json:"nodeType"`
	GatewayId       int64  `json:"gatewayId"` //网关ID
	DriveInstanceId string `json:"driveInstanceId"`
	Name            string `json:"name"binding:"required" `
	Status          string `json:"status"`
	Sn              string `json:"sn"`
	DevNo           string `json:"devNo"`
	Description     string `json:"description"`
	Secret          string `json:"secret"`
	InstallLocation string `json:"installLocation"`
	GatewayName     string `json:"gatewayName"`
}
