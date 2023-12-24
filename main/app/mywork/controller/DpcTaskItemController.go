//	=========================================================================
//
// LV自动生成控制器相关代码，只生成一次，按需修改,默认再次生成不会覆盖.
// date：2023-12-24 16:29:05 +0800 CST
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
type DpcTaskItemController struct{}

// List 查询页
func (w DpcTaskItemController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "mywork/item/list.html").WriteTpl()
}
// Add 新增页
func (w DpcTaskItemController) Add(c *gin.Context) {
	lv_web.BuildTpl(c, "mywork/item/add.html").WriteTpl()
}
// Edit 修改页
func (w DpcTaskItemController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
    itemService :=service.DpcTaskItemService{}
	entity, err := itemService.FindById(id)
	if err != nil || entity == nil {
		lv_web.ErrorTpl(c).WriteTpl(gin.H{"desc": "数据不存在",})
		return
	}
	lv_web.BuildTpl(c, "mywork/item/edit.html").WriteTpl(gin.H{
		"item": entity,
	})
}

//	========================================================================
//
// api
// =========================================================================

// ListAjax 新增页面保存
func (w DpcTaskItemController) ListAjax(c *gin.Context) {
	req := new(vo.PageDpcTaskItemReq)
	err := c.ShouldBind(req)
	lv_logic.HasErrAndPanic(err)
	var svc service.DpcTaskItemService
	result, total, _ := svc.ListByPage(req)
	lv_web.SucessPage(c, result, total)
}

// AddSave 新增页面保存
func (w DpcTaskItemController) AddSave(c *gin.Context) {
	req := new(vo.AddDpcTaskItemReq)
    err := c.ShouldBind(req)
    lv_logic.HasErrAndPanic(err)
	var svc service.DpcTaskItemService

    var userService sysService.UserService
    user := userService.GetProfile(c)
	id, err := svc.AddSave(req, user)
    lv_logic.HasErrAndPanic(err)
	lv_web.SucessData(c, id)
}

// EditSave 修改页面保存
func (w DpcTaskItemController) EditSave(c *gin.Context) {
	req := new(vo.EditDpcTaskItemReq)
	err := c.ShouldBind(req)
    lv_logic.HasErrAndPanic(err)
    var svc service.DpcTaskItemService

    var userService sysService.UserService
    user := userService.GetProfile(c)
	err = svc.EditSave(req, user)
	lv_logic.HasErrAndPanic(err)
	lv_web.Success(c, nil, "success")
}

// Remove 删除数据
func (w DpcTaskItemController) Remove(c *gin.Context) {
	req := new(model_cmn.RemoveReq)
	err := c.ShouldBind(req)
    lv_logic.HasErrAndPanic(err)
	var svc service.DpcTaskItemService
	rs := svc.DeleteByIds(req.Ids)
	lv_web.SuccessData(c, rs)
}

// 导出
func (w DpcTaskItemController) Export(c *gin.Context) {
	req := new(vo.PageDpcTaskItemReq)
	err := c.ShouldBind(req)
    lv_logic.HasErrAndPanic(err)
	var svc service.DpcTaskItemService
	url, err := svc.ExportAll(req)
	lv_logic.HasErrAndPanic(err)
	lv_web.SucessDataMsg(c, url, url)
}
