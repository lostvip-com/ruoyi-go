package controller

import (
	util2 "common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_logic"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	postService "main/internal/system/service"
	"main/internal/system/vo"
)

type PostController struct {
}

// 列表页
func (w *PostController) List(c *gin.Context) {
	util2.BuildTpl(c, "system/post/list").WriteTpl()
}

// 列表分页数据
func (w *PostController) ListAjax(c *gin.Context) {
	var req *vo.SelectPostPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util2.ErrorResp(c).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
		return
	}
	var postService postService.SysPostService
	result, total, err := postService.SelectListByPage(req)
	if err != nil {
		panic(err)
	}
	util2.BuildTable(c, total, result).WriteJsonExit()
}

// 新增页面
func (w *PostController) Add(c *gin.Context) {
	util2.BuildTpl(c, "system/post/add").WriteTpl()
}

// 新增页面保存
func (w *PostController) AddSave(c *gin.Context) {
	var req *vo.AddPostReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util2.ErrorResp(c).SetBtype(lv_dto.Buniss_Add).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
		return
	}
	var postService postService.SysPostService
	if postService.IsPostCodeExist(req.PostCode) {
		util2.Fail(c, "岗位编码已存在")
		return
	}
	pid, err := postService.AddSave(req, c)
	if err != nil || pid <= 0 {
		util2.Fail(c, "添加失败！")
		return
	}
	util2.Success(c, pid, "添加成功！")
}

// 修改页面
func (w *PostController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	var postService postService.SysPostService
	post, err := postService.SelectRecordById(id)
	if err != nil || post == nil {
		util2.WriteErrorTPL(c, gin.H{
			"desc": "岗位不存在",
		})
		return
	}
	util2.WriteTpl(c, "system/post/edit", gin.H{
		"post": post,
	})
}

// EditSave 修改页面保存
func (w *PostController) EditSave(c *gin.Context) {
	var req *vo.EditSysPostReq
	if err := c.ShouldBind(&req); err != nil {
		util2.ErrorResp(c).SetBtype(lv_dto.Buniss_Edit).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
		return
	}
	var postService postService.SysPostService
	err := postService.EditSave(req, c)
	if err != nil {
		util2.ErrorResp(c).SetBtype(lv_dto.Buniss_Edit).Log("岗位管理", req).WriteJsonExit()
		return
	}
	util2.SucessResp(c).SetBtype(lv_dto.Buniss_Edit).Log("岗位管理", req).WriteJsonExit()
}

// Remove 删除数据
func (w *PostController) Remove(c *gin.Context) {
	var req *lv_dto.IdsReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util2.ErrorResp(c).SetMsg(err.Error()).SetBtype(lv_dto.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
		return
	}
	var postService postService.SysPostService
	rs := postService.DeleteRecordByIds(req.Ids)
	if rs > 0 {
		util2.SucessResp(c).SetBtype(lv_dto.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
	} else {
		util2.ErrorResp(c).SetBtype(lv_dto.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
	}
}

// 导出
func (w *PostController) Export(c *gin.Context) {
	var req *vo.SelectPostPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util2.ErrorResp(c).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
		return
	}
	var postService postService.SysPostService
	url, err := postService.Export(req)

	if err != nil {
		util2.ErrorResp(c).SetMsg(err.Error()).Log("岗位管理", req).WriteJsonExit()
		return
	}
	util2.SucessResp(c).SetMsg(url).SetBtype(lv_dto.Buniss_Del).Log("岗位管理", req).WriteJsonExit()
}

// 检查岗位编码是否已经存在
func (w *PostController) IsPostCodeExist(c *gin.Context) {
	var req *vo.CheckPostCodeALLReq
	if err := c.ShouldBind(&req); err != nil {
		util2.Fail(c, "参数错误！")
		return
	}
	var postService postService.SysPostService
	exist := postService.IsPostCodeExist(req.PostCode)
	util2.Success(c, exist, lv_logic.IfTrue(exist, "编码冲突！", "ok").(string))
}
