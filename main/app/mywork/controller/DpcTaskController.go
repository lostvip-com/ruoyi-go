//	=========================================================================
//
// LV自动生成控制器相关代码，只生成一次，按需修改,默认再次生成不会覆盖.
// date：2023-12-24 14:57:01 +0800 CST
// author：lv
// ==========================================================================
package controller

import (
    "github.com/gin-gonic/gin"
    "lostvip.com/utils/lv_conv"
    "lostvip.com/utils/lv_web"
    "lostvip.com/utils/lv_logic"
    "robvi/app/common/model_cmn"
    "robvi/app/mywork/vo"
    "robvi/app/mywork/service"
    sysService "robvi/app/system/service"
)
type DpcTaskController struct{}

// List 查询页
func (w DpcTaskController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "mywork/task/list.html").WriteTpl()
}
// Add 新增页
func (w DpcTaskController) Add(c *gin.Context) {
	lv_web.BuildTpl(c, "mywork/task/add.html").WriteTpl()
}
// Edit 修改页
func (w DpcTaskController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
    taskService :=service.DpcTaskService{}
	entity, err := taskService.FindById(id)
	if err != nil || entity == nil {
		lv_web.ErrorTpl(c).WriteTpl(gin.H{"desc": "数据不存在",})
		return
	}
	lv_web.BuildTpl(c, "mywork/task/edit.html").WriteTpl(gin.H{
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
	lv_logic.HasErrAndPanic(err)
	var svc service.DpcTaskService
	result, total, _ := svc.ListByPage(req)
	lv_web.SucessPage(c, result, total)
}

// AddSave 新增页面保存
func (w DpcTaskController) AddSave(c *gin.Context) {
	req := new(vo.AddDpcTaskReq)
    err := c.ShouldBind(req)
    lv_logic.HasErrAndPanic(err)
	var svc service.DpcTaskService

    var userService sysService.UserService
    user := userService.GetProfile(c)
    req.CreateBy = user.LoginName
	id, err := svc.AddSave(req)
    lv_logic.HasErrAndPanic(err)
	lv_web.SucessData(c, id)
}

// EditSave 修改页面保存
func (w DpcTaskController) EditSave(c *gin.Context) {
	req := new(vo.EditDpcTaskReq)
	err := c.ShouldBind(req)
    lv_logic.HasErrAndPanic(err)
    var svc service.DpcTaskService
    var userService sysService.UserService
    user := userService.GetProfile(c)
    req.UpdateBy = user.LoginName
	err = svc.EditSave(req)
	lv_logic.HasErrAndPanic(err)
	lv_web.Success(c, nil, "success")
}

// Remove 删除数据
func (w DpcTaskController) Remove(c *gin.Context) {
	req := new(model_cmn.RemoveReq)
	err := c.ShouldBind(req)
    lv_logic.HasErrAndPanic(err)
	var svc service.DpcTaskService
	rs := svc.DeleteByIds(req.Ids)
	lv_web.SuccessData(c, rs)
}

// 导出
func (w DpcTaskController) Export(c *gin.Context) {
	req := new(vo.PageDpcTaskReq)
	err := c.ShouldBind(req)
    lv_logic.HasErrAndPanic(err)
	var svc service.DpcTaskService
	url, err := svc.ExportAll(req)
	lv_logic.HasErrAndPanic(err)
	lv_web.SucessDataMsg(c, url, url)
}
