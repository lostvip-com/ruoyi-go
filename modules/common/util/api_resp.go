package util

import (
	"common/session"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"net/http"
)

// 返回一个成功的消息体
func SucessData(c *gin.Context, data any) {
	msg := lv_dto.CommonRes{
		Code: 200,
		Data: data,
		Msg:  "操作成功",
	}
	c.AbortWithStatusJSON(http.StatusOK, &msg)
}

// 返回一个成功的消息体
func SucessDataMsg(c *gin.Context, data any, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, &lv_dto.CommonRes{
		Code: 200,
		Data: data,
		Msg:  msg,
	})
}

// SucessPage 返回一个成功的消息体
func SucessPage(c *gin.Context, rows any, total int64) {
	msg := lv_dto.TableDataInfo{
		Code:  200,
		Rows:  rows,
		Total: total,
		Msg:   "Success!",
	}
	c.AbortWithStatusJSON(http.StatusOK, &msg)
}

// 返回一个成功的消息体
func Fail(c *gin.Context, msg string) {
	ret := lv_dto.CommonRes{
		Code: 500,
		Msg:  msg,
	}
	c.AbortWithStatusJSON(http.StatusOK, &ret)
}
func WriteTpl(c *gin.Context, tpl string, params ...gin.H) {
	var data gin.H
	login := session.GetLoginInfo(c)
	if params == nil || len(params) == 0 {
		data = gin.H{}
	} else {
		data = params[0]
	}
	for k, v := range c.Keys {
		data[k] = v
	}
	data["uid"] = login.UserId
	c.HTML(http.StatusOK, tpl, data)
	c.Abort()
}

// WriteErrorTPL 返回一个错误的tpl响应
func WriteErrorTPL(c *gin.Context, params ...gin.H) {
	WriteTpl(c, "error/error.html", params...)
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////
func Err(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 500, "msg": msg, "data": ""})
}

// 通常成功数据处理
func Success(c *gin.Context, data interface{}, msg string) {
	//如果传进来的是指针就对其进行翻译，不是指针保持原来的内容不变
	//if data!=nil{
	//	if reflect.TypeOf(data).Kind() == reflect.Ptr {
	//		util.TranslateByTag(data)
	//	}
	//}
	//msg = global.GetTextLocale(msg)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 200, "msg": msg, "data": data})
}

// 通常成功数据处理
func SuccessData(c *gin.Context, data interface{}) {
	//如果传进来的是指针就对其进行翻译，不是指针保持原来的内容不变
	//if data!=nil{
	//	if reflect.TypeOf(data).Kind() == reflect.Ptr {
	//		util.TranslateByTag(data)
	//	}
	//}
	//msg = global.GetTextLocale(msg)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": data})
}

// 失败数据处理
func Error(c *gin.Context, err error) {
	var msg string
	if err != nil {
		msg = err.Error()
	}
	Err(c, msg)
}
func ErrResp(c *gin.Context, res lv_dto.Resp) {
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// PageOK 分页数据处理 ， 自动翻译 Tag locale标记的字段
func PageOK(c *gin.Context, result lv_dto.RespPage) {
	c.AbortWithStatusJSON(http.StatusOK, result)
}

// PageOK2 分页数据处理 ， 自动翻译 Tag locale标记的字段
func PageOK2(c *gin.Context, rows any, total int64) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "rows": rows, "total": total})
}

///**
// * 自动翻译 Tag locale标记的字段
// */
//func PageOKLocale(c *gin.Context, result []interface{}, count int, pageIndex int, pageSize int, msg string) {
//	var res Page
//	util.TranslateSliceByTag(result)
//	// result.([]interface{})
//	res.List = result
//	res.Count = count
//	res.PageIndex = pageIndex
//	res.PageSize = pageSize
//	Ok(c, res, msg)
//}

// 兼容函数
func Custum(c *gin.Context, data gin.H) {
	c.AbortWithStatusJSON(http.StatusOK, data)
}
