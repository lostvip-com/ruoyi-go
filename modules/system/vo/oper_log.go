package vo

import "github.com/lostvip-com/lv_framework/web/lv_dto"

// 查询列表请求参数
type OperLogPageReq struct {
	Title         string `form:"title"`         //系统模块
	OperName      string `form:"operName"`      //操作人员
	BusinessTypes int    `form:"businessTypes"` //操作类型
	Status        string `form:"status"`        //操作类型
	BeginTime     string `form:"beginTime"`     //数据范围
	EndTime       string `form:"endTime"`       //开始时间
	lv_dto.Paging
}
