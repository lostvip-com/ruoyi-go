// ==========================================================================
// LV自动生成model扩展代码列表、增、删，改、查、导出，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2020-02-17 14:03:51
// 生成路径: app/model_cmn/online/online.go
// 生成人：yunjie
// ==========================================================================
package vo

import (
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"time"
)

// 新增页面请求参数
type AddOnlineReq struct {
	LoginName      string    `form:"loginName" binding:"required"`
	DeptName       string    `form:"deptName" binding:"required"`
	Ipaddr         string    `form:"ipaddr" `
	LoginLocation  string    `form:"loginLocation" `
	Browser        string    `form:"browser" `
	Os             string    `form:"os" `
	Status         string    `form:"status" binding:"required"`
	StartTimestamp time.Time `form:"startTimestamp" `
	LastAccessTime time.Time `form:"lastAccessTime" `
	ExpireTime     int       `form:"expireTime" `
}

// 分页请求参数
type OnlinePageReq struct {
	SessionId      string    `form:"sessionId"`      //用户会话id
	LoginName      string    `form:"loginName"`      //登录账号
	DeptName       string    `form:"deptName"`       //部门名称
	Ipaddr         string    `form:"ipaddr"`         //登录IP地址
	LoginLocation  string    `form:"loginLocation"`  //登录地点
	Browser        string    `form:"browser"`        //浏览器类型
	Os             string    `form:"os"`             //操作系统
	Status         string    `form:"status"`         //在线状态on_line在线off_line离线
	StartTimestamp time.Time `form:"startTimestamp"` //session创建时间
	LastAccessTime time.Time `form:"lastAccessTime"` //session最后访问时间
	ExpireTime     int       `form:"expireTime"`     //超时时间，单位为分钟
	BeginTime      string    `form:"beginTime"`      //开始时间
	EndTime        string    `form:"endTime"`        //结束时间
	lv_dto.Paging
}
