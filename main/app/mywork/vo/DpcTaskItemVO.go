// ==========================================================================
// LV自动生成model扩展代码列表 按需修改
// 生成日期：2023-12-24 16:29:05 +0800 CST
// 生成人：lv
// ==========================================================================
package vo

import (
)

//新增页面请求参数
type AddDpcTaskItemReq struct { 
	 
	 TaskId  int64   `form:"taskId"  `  
	 Action  string   `form:"action"  `  
	 Content  string   `form:"content"  `  
	 IdxXpath  string   `form:"idxXpath"  `  
	 IdxImg1  string   `form:"idxImg1"  `  
	 IdxImg2  string   `form:"idxImg2"  `  
	 IdxImg3  string   `form:"idxImg3"  `  
	 Status  string   `form:"status" binding:"required" `  
	 Sort  int64   `form:"sort"  `  
	 DelFlag  string   `form:"delFlag"  `  
	 
	 
	 
	 
}

//修改页面请求参数
type EditDpcTaskItemReq struct {
      Id    int64  `form:"id" binding:"required"`    
      TaskId  int64 `form:"taskId"  `   
      Action  string `form:"action"  `   
      Content  string `form:"content"  `   
      IdxXpath  string `form:"idxXpath"  `   
      IdxImg1  string `form:"idxImg1"  `   
      IdxImg2  string `form:"idxImg2"  `   
      IdxImg3  string `form:"idxImg3"  `   
      Status  string `form:"status" binding:"required" `   
      Sort  int64 `form:"sort"  `            
}

//分页请求参数 
type PageDpcTaskItemReq struct {



TaskId int64 `form:"taskId"` //任务ID

Action string `form:"action"` //click dbclick write write_enter

Content string `form:"content"` //内容

IdxXpath string `form:"idxXpath"` //html xpath





















	BeginTime  string `form:"beginTime"`  //开始时间
	EndTime    string `form:"endTime"`    //结束时间
	PageNum    int    `form:"pageNum"`    //当前页码
	PageSize   int    `form:"pageSize"`   //每页数
}
