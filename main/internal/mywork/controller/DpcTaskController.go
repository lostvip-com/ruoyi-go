//	=========================================================================
//
// LV自动生成控制器相关代码，只生成一次，按需修改,默认再次生成不会覆盖.
// date：2024-01-03 21:50:54 +0800 CST
// author：lv
// ==========================================================================
package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"main/internal/mywork/service"
	"main/internal/mywork/vo"
	sysService "main/internal/system/service"
)

type DpcTaskController struct{}

// List 查询页
func (w DpcTaskController) List(c *gin.Context) {
	util.WriteTpl(c, "mywork/task/list.html")
}

// Add 新增页
func (w DpcTaskController) Add(c *gin.Context) {
	util.BuildTpl(c, "mywork/task/add.html").WriteTpl()
}

// Edit 修改页
func (w DpcTaskController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	taskService := service.DpcTaskService{}
	entity, err := taskService.FindById(id)
	if err != nil || entity == nil {
		util.ErrorTpl(c).WriteTpl(gin.H{"desc": "数据不存在"})
		return
	}
	util.BuildTpl(c, "mywork/task/edit.html").WriteTpl(gin.H{
		"task": entity,
	})
}

//	========================================================================
//
// api
// =========================================================================

// ListAjax 新增页面保存
func (w DpcTaskController) ListAjax(c *gin.Context) {
	req := new(vo.PageDpcTaskReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.DpcTaskService
	result, total, _ := svc.ListByPage(req)
	util.SucessPage(c, result, total)
}

// AddSave 新增页面保存
func (w DpcTaskController) AddSave(c *gin.Context) {
	req := new(vo.AddDpcTaskReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.DpcTaskService

	var userService sysService.UserService
	user := userService.GetProfile(c)
	req.CreateBy = user.LoginName
	id, err := svc.AddSave(req)
	lv_err.HasErrAndPanic(err)
	util.SucessData(c, id)
}

// EditSave 修改页面保存
func (w DpcTaskController) EditSave(c *gin.Context) {
	req := new(vo.EditDpcTaskReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.DpcTaskService
	var userService sysService.UserService
	user := userService.GetProfile(c)
	req.UpdateBy = user.LoginName
	err = svc.EditSave(req)
	lv_err.HasErrAndPanic(err)
	util.Success(c, nil, "success")
}

// Remove 删除数据
func (w DpcTaskController) Remove(c *gin.Context) {
	req := new(lv_dto.IdsReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.DpcTaskService
	rs := svc.DeleteByIds(req.Ids)
	util.SuccessData(c, rs)
}

// 导出
func (w DpcTaskController) Export(c *gin.Context) {
	req := new(vo.PageDpcTaskReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.DpcTaskService
	url, err := svc.ExportAll(req)
	lv_err.HasErrAndPanic(err)
	util.SucessDataMsg(c, url, url)
}
