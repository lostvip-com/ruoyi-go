package lv_web

import (
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/web/dto"
	"net/http"
)

func HTML(c *gin.Context, path string, data gin.H) {
	for k, v := range c.Keys {
		data[k] = v
	}
	c.HTML(http.StatusOK, path, data)
}

func Err(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 500, "msg": msg, "data": ""})
}

func PageOKOld(c *gin.Context, result interface{}, total int, msg string) {
	//util.TranslateSliceByTag(slice)
	data := map[string]interface{}{"list": result, "total": total}
	Success(c, data, msg)
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
func ErrResp(c *gin.Context, res dto.Resp) {
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// 分页数据处理 ， 自动翻译 Tag locale标记的字段
func PageOK(c *gin.Context, result dto.RespPage) {
	c.AbortWithStatusJSON(http.StatusOK, result)
}

// 分页数据处理 ， 自动翻译 Tag locale标记的字段
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
