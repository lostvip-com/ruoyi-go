package vo

import "lostvip.com/utils/lv_web"

// Fill with you ideas below.
// 新增页面请求参数
type AddConfigReq struct {
	ConfigName  string `form:"configName"  binding:"required"`
	ConfigKey   string `form:"configKey"  binding:"required"`
	ConfigValue string `form:"configValue"  binding:"required"`
	ConfigType  string `form:"configType"    binding:"required"`
	Remark      string `form:"remark"`
}

// 修改页面请求参数
type EditConfigReq struct {
	ConfigId    int64  `form:"configId" binding:"required"`
	ConfigName  string `form:"configName"  binding:"required"`
	ConfigKey   string `form:"configKey"  binding:"required"`
	ConfigValue string `form:"configValue"  binding:"required"`
	ConfigType  string `form:"configType"    binding:"required"`
	Remark      string `form:"remark"`
}

// 分页请求参数
type SelectConfigPageReq struct {
	ConfigName string `form:"configName"` //参数名称
	ConfigKey  string `form:"configKey"`  //参数键名
	ConfigType string `form:"configType"` //状态
	BeginTime  string `form:"beginTime"`  //开始时间
	EndTime    string `form:"endTime"`    //结束时间
	lv_web.Paging
}

// 检查参数键名请求参数
type CheckConfigKeyReq struct {
	ConfigId  int64  `form:"configId"  binding:"required"`
	ConfigKey string `form:"configKey"  binding:"required"`
}

//
//// 检查参数键名请求参数
//type CheckPostCodeALLReq struct {
//	ConfigKey string `form:"configKey"  binding:"required"`
//}
