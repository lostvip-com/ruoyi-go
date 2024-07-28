package dto

type BunissType int

// 业务类型
const (
	Buniss_Other BunissType = 0 //0其它
	Buniss_Add   BunissType = 1 //1新增
	Buniss_Edit  BunissType = 2 //2修改
	Buniss_Del   BunissType = 3 //3删除
)

// 响应结果
const (
	SUCCESS      = 200 // 成功
	ERROR        = 500 //错误
	UNAUTHORIZED = 403 //无权限
	FAIL         = -1  //失败
)

// 错误处理页面
const (
	ERROR_PAGE  = "error/error.html"  //错误提示页面
	UNAUTH_PAGE = "error/unauth.html" //无权限提示页面
)

// 通用api响应
type CommonRes struct {
	Code  int         `json:"code"`  //响应编码 0 成功 500 错误 403 无权限  -1  失败
	Msg   string      `json:"msg"`   //消息
	Data  interface{} `json:"data"`  //数据内容
	Btype BunissType  `json:"otype"` //业务类型
}

// 验证码响应
type CaptchaRes struct {
	Code           int         `json:"code"` //响应编码 0 成功 500 错误 403 无权限
	Msg            string      `json:"msg"`  //消息
	Img            interface{} `json:"img"`  //数据内容
	Uuid           string      `json:"uuid"` //验证码ID
	CaptchaEnabled bool        `json:"captchaEnabled"`
}

// 通用分页表格响应
type TableDataInfo struct {
	Total any         `json:"total"` //总数
	Rows  interface{} `json:"rows"`  //数据
	Code  int         `json:"code"`  //响应编码 200 成功 500 错误 403 无权限
	Msg   string      `json:"msg"`   //消息
}

// 通用的树形结构
type Ztree struct {
	Id       int64  `json:"id"`       //节点ID
	Pid      int64  `json:"pId"`      //节点父ID
	Name     string `json:"name"`     //节点名称
	Title    string `json:"title"`    //节点标题
	Checked  bool   `json:"checked"`  //是否勾选
	Open     bool   `json:"open"`     //是否展开
	Nocheck  bool   `json:"nocheck"`  //是否能勾选
	NodeType string `json:"nodeType"` //节点类型
}

// 通用的删除请求
type IdsReq struct {
	Ids string `form:"ids"  binding:"required"`
}

// 通用详情请求
type DetailReq struct {
	Id int64 `json:"id"` //主键ID
}

// 通用修改请求
type EditReq struct {
	Id int64 `json:"id"` //主键ID
}
