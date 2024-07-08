package vo

// Fill with you ideas below.
// 新增页面请求参数
type AddDictDataReq struct {
	DictLabel string `form:"dictLabel"  binding:"required"`
	DictValue string `form:"dictValue"  binding:"required"`
	DictType  string `form:"dictType"  binding:"required"`
	DictSort  int    `form:"dictSort"  binding:"required"`
	CssClass  string `form:"cssClass"`
	ListClass string `form:"listClass" binding:"required"`
	IsDefault string `form:"isDefault" binding:"required"`
	Status    string `form:"status"    binding:"required"`
	Remark    string `form:"remark"`
}

// 修改页面请求参数
type EditDictDataReq struct {
	DictCode  int64  `form:"dictCode" binding:"required"`
	DictLabel string `form:"dictLabel"  binding:"required"`
	DictValue string `form:"dictValue"  binding:"required"`
	DictType  string `form:"dictType"`
	DictSort  int    `form:"dictSort"  binding:"required"`
	CssClass  string `form:"cssClass"`
	ListClass string `form:"listClass" binding:"required"`
	IsDefault string `form:"isDefault" binding:"required"`
	Status    string `form:"status"    binding:"required"`
	Remark    string `form:"remark"`
}

// 分页请求参数
type SelectDictDataPageReq struct {
	DictType  string `form:"dictType"`  //字典名称
	DictLabel string `form:"dictLabel"` //字典标签
	Status    string `form:"status"`    //状态
	BeginTime string `form:"beginTime"` //开始时间
	EndTime   string `form:"endTime"`   //结束时间
	PageNum   int    `form:"pageNum"`   //当前页码
	PageSize  int    `form:"pageSize"`  //每页数
}
