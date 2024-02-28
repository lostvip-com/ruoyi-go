//	=========================================================================
//
// LV自动生成控制器相关代码，只生成一次，按需修改,默认再次生成不会覆盖.
// date：2024-02-28 14:21:50 +0800 CST
// author：lv
// ==========================================================================
package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_err"
	"lostvip.com/utils/lv_web"
	"lostvip.com/web/dto"
	"main/app/mywork/service"
	"main/app/mywork/vo"
	sysService "main/app/system/service"
)

type AppGenParamsController struct{}

//	========================================================================
//
// api
// =========================================================================

// ListAjax 新增页面保存
func (w AppGenParamsController) ListAjax(c *gin.Context) {
	req := new(vo.PageAppGenParamsReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.AppGenParamsService
	result, total, _ := svc.ListByPage(req)
	lv_web.SucessPage(c, result, total)
}

// AddSave 新增页面保存
func (w AppGenParamsController) AddSave(c *gin.Context) {
	req := new(vo.AddAppGenParamsReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.AppGenParamsService

	var userService sysService.UserService
	user := userService.GetProfile(c)
	req.CreateBy = user.LoginName
	id, err := svc.AddSave(req)
	lv_err.HasErrAndPanic(err)
	lv_web.SucessData(c, id)
}

// EditSave 修改页面保存
func (w AppGenParamsController) EditSave(c *gin.Context) {
	req := new(vo.EditAppGenParamsReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.AppGenParamsService
	var userService sysService.UserService
	user := userService.GetProfile(c)
	req.UpdateBy = user.LoginName
	err = svc.EditSave(req)
	lv_err.HasErrAndPanic(err)
	lv_web.Success(c, nil, "success")
}

// Remove 删除数据
func (w AppGenParamsController) Remove(c *gin.Context) {
	req := new(dto.RemoveReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.AppGenParamsService
	rs := svc.DeleteByIds(req.Ids)
	lv_web.SuccessData(c, rs)
}

// 导出
func (w AppGenParamsController) Export(c *gin.Context) {
	req := new(vo.PageAppGenParamsReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.AppGenParamsService
	url, err := svc.ExportAll(req)
	lv_err.HasErrAndPanic(err)
	lv_web.SucessDataMsg(c, url, url)
}
