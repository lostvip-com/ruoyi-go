// ==========================================================================
// LV自动生成model扩展代码列表 按需修改
// 生成日期：2024-01-02 15:41:17 +0800 CST
// 生成人：lv
// ==========================================================================
package vo

import (
"time"
)

//新增页面请求参数
type AddDpcTableReq struct {
	Username string `form:"username" binding:"required" `
	Password string `form:"password"  `
	TaskContent string `form:"taskContent"  `
	StartDate time.Time `form:"startDate"  `
	EndDate time.Time `form:"endDate"  `
	WorkDays int64 `form:"workDays"  `
	AutoSubmit int `form:"autoSubmit"  `
	Status string `form:"status" binding:"required" `
	Sort int64 `form:"sort"  `
	UpdateTime time.Time `form:"updateTime"  `
    CreateBy string
}

//修改页面请求参数
type EditDpcTableReq struct {
      Id    int64  `form:"id" binding:"required"`    
      Username  string `form:"username" binding:"required" `   
      Password  string `form:"password"  `   
      TaskContent  string `form:"taskContent"  `   
      StartDate  time.Time `form:"startDate"  `   
      EndDate  time.Time `form:"endDate"  `   
      WorkDays  int64 `form:"workDays"  `   
      AutoSubmit  int `form:"autoSubmit"  `   
      Status  string `form:"status" binding:"required" `   
      Sort  int64 `form:"sort"  `   
      UpdateTime  time.Time `form:"updateTime"  `  
      UpdateBy string
}

//分页请求参数 
type PageDpcTableReq struct {
	Username string `form:"username"` //
	Password string `form:"password"` //
	TaskContent string `form:"taskContent"` //
	StartDate time.Time `form:"startDate"` //
	BeginTime  string `form:"beginTime"`  //开始时间
	EndTime    string `form:"endTime"`    //结束时间
	PageNum    int    `form:"pageNum"`    //当前页码
	PageSize   int    `form:"pageSize"`   //每页数
}
