package post

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_web"
	"robvi/app/sys/model"
	"robvi/app/sys/model/system/post"
	postService "robvi/app/sys/service/system/post"
)

// 列表页
func List(c *gin.Context) {
	lv_web.BuildTpl(c, "system/post/list").WriteTpl()
}

// 列表分页数据
func ListAjax(c *gin.Context) {
	var req *post.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
		return
	}
	rows := make([]post.SysPost, 0)
	result, page, err := postService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	lv_web.BuildTable(c, page.Total, rows).WriteJsonExit()
}

// 新增页面
func Add(c *gin.Context) {
	lv_web.BuildTpl(c, "system/post/add").WriteTpl()
}

// 新增页面保存
func AddSave(c *gin.Context) {
	var req *post.AddReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
		return
	}

	if postService.CheckPostNameUniqueAll(req.PostName) == "1" {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("岗位名称已存在").Log("岗位管理", req).WriteJsonExit()
		return
	}

	if postService.CheckPostCodeUniqueAll(req.PostCode) == "1" {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("岗位编码已存在").Log("岗位管理", req).WriteJsonExit()
		return
	}

	pid, err := postService.AddSave(req, c)

	if err != nil || pid <= 0 {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).Log("岗位管理", req).WriteJsonExit()
		return
	}
	lv_web.ErrorResp(c).SetData(pid).SetBtype(model.Buniss_Add).Log("岗位管理", req).WriteJsonExit()
}

// 修改页面
func Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}

	post, err := postService.SelectRecordById(id)

	if err != nil || post == nil {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "岗位不存在",
		})
		return
	}

	lv_web.BuildTpl(c, "system/post/edit").WriteTpl(gin.H{
		"post": post,
	})
}

// 修改页面保存
func EditSave(c *gin.Context) {
	var req *post.EditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
		return
	}

	if postService.CheckPostNameUnique(req.PostName, req.PostId) == "1" {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg("岗位名称已存在").Log("岗位管理", req).WriteJsonExit()
		return
	}

	if postService.CheckPostCodeUnique(req.PostCode, req.PostId) == "1" {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg("岗位编码已存在").Log("岗位管理", req).WriteJsonExit()
		return
	}

	rs, err := postService.EditSave(req, c)

	if err != nil || rs <= 0 {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).Log("岗位管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetData(rs).SetBtype(model.Buniss_Edit).Log("岗位管理", req).WriteJsonExit()
}

// 删除数据
func Remove(c *gin.Context) {
	var req *model.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).SetBtype(model.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
		return
	}

	rs := postService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		lv_web.SucessResp(c).SetBtype(model.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
	}
}

// 导出
func Export(c *gin.Context) {
	var req *post.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
		return
	}
	url, err := postService.Export(req)

	if err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetMsg(url).SetBtype(model.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
}

// 检查岗位名称是否已经存在不包括本岗位
func CheckPostNameUnique(c *gin.Context) {
	var req *post.CheckPostNameReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := postService.CheckPostNameUnique(req.PostName, req.PostId)

	c.Writer.WriteString(result)
}

// 检查岗位名称是否已经存在
func CheckPostNameUniqueAll(c *gin.Context) {
	var req *post.CheckPostNameALLReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := postService.CheckPostNameUniqueAll(req.PostName)

	c.Writer.WriteString(result)
}

// 检查岗位编码是否已经存在不包括本岗位
func CheckPostCodeUnique(c *gin.Context) {
	var req *post.CheckPostCodeReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := postService.CheckPostCodeUnique(req.PostCode, req.PostId)

	c.Writer.WriteString(result)
}

// 检查岗位编码是否已经存在
func CheckPostCodeUniqueAll(c *gin.Context) {
	var req *post.CheckPostCodeALLReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := postService.CheckPostCodeUniqueAll(req.PostCode)

	c.Writer.WriteString(result)
}
