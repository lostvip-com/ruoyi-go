package vo

import (
	"github.com/lostvip-com/lv_framework/web/lv_dto"
)

// Fill with you ideas below.
// 查询列表请求参数
type LoginInfoPageReq struct {
	LoginName string `form:"loginName"` //登录名
	Status    string `form:"status"`    //状态
	Ipaddr    string `form:"ipaddr"`    //登录地址
	BeginTime string `form:"beginTime"` //数据范围
	EndTime   string `form:"endTime"`   //开始时间
	lv_dto.Paging
}
