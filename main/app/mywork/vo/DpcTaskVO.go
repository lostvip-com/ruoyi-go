// ==========================================================================
// LV自动生成model扩展代码列表 按需修改
// 生成日期：2024-01-03 21:50:54 +0800 CST
// 生成人：lv
// ==========================================================================
package vo

import (
"time"
)

//新增页面请求参数
type AddDpcTaskReq struct {
	Username string `form:"username"binding:"required" `
	Password string `form:"password"`
	PrjCode string `form:"prjCode"`
	TaskContent string `form:"taskContent"`
	StartDate time.Time `form:"startDate" time_format:"2006-01-02" `
	EndDate time.Time `form:"endDate" time_format:"2006-01-02" `
	WorkDays int64 `form:"workDays"`
	AutoSubmit string `form:"autoSubmit"`
	Status string `form:"status"binding:"required" `
	Sort int64 `form:"sort"`
	DelFlag int `form:"delFlag"`
    CreateBy string
}

//修改页面请求参数
type EditDpcTaskReq struct {
      Id    int64  `form:"id" binding:"required"`
            Username string `form:"username"binding:"required" `
            Password string `form:"password"`
            PrjCode string `form:"prjCode"`
            TaskContent string `form:"taskContent"`
            StartDate time.Time `form:"startDate" time_format:"2006-01-02" `
            EndDate time.Time `form:"endDate" time_format:"2006-01-02" `
            WorkDays int64 `form:"workDays"`
            AutoSubmit string `form:"autoSubmit"`
            Status string `form:"status"binding:"required" `
            Sort int64 `form:"sort"`
            DelFlag int `form:"delFlag"`
      UpdateBy string
}

//分页请求参数 
type PageDpcTaskReq struct {
	Username string `form:"username"` //工号
	Password string `form:"password"` //密码
	PrjCode string `form:"prjCode"` //项  目  号
	TaskContent string `form:"taskContent"` //任务内容
	BeginTime  string `form:"beginTime"`  //开始时间
	EndTime    string `form:"endTime"`    //结束时间
	PageNum    int    `form:"pageNum"`    //当前页码
	PageSize   int    `form:"pageSize"`   //每页数
}
