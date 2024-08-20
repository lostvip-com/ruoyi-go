//	=========================================================================
//
// LV自动生成控制器相关代码，只生成一次，按需修改,默认再次生成不会覆盖.
// date：2024-07-30 21:59:34 +0800 CST
// author：lv
// ==========================================================================
package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"main/internal/iot_dev/service"
	"main/internal/iot_dev/vo"
	sysService "main/internal/system/service"
)

type IotPrdActionController struct{}

//	========================================================================
//
// api
// =========================================================================

// ListAjax 新增页面保存
func (w IotPrdActionController) ListAjax(c *gin.Context) {
	req := new(vo.PageIotPrdActionReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	if req.PageSize == 0 {
		req.PageSize = lv_dto.PageSize
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	var svc service.IotPrdActionService
	result, total, _ := svc.ListByPage(req)
	util.SucessPage(c, result, total)
}

// AddSave 新增页面保存
func (w IotPrdActionController) AddSave(c *gin.Context) {
	req := new(vo.AddIotPrdActionReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.IotPrdActionService

	var userService sysService.UserService
	user := userService.GetProfile(c)
	req.CreateBy = user.LoginName
	id, err := svc.AddSave(req)
	lv_err.HasErrAndPanic(err)
	util.SucessData(c, id)
}

// EditSave 修改页面保存
func (w IotPrdActionController) EditSave(c *gin.Context) {
	req := new(vo.EditIotPrdActionReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.IotPrdActionService
	var userService sysService.UserService
	user := userService.GetProfile(c)
	req.UpdateBy = user.LoginName
	err = svc.EditSave(req)
	lv_err.HasErrAndPanic(err)
	util.Success(c, nil, "success")
}

// Remove 删除数据
func (w IotPrdActionController) Remove(c *gin.Context) {
	req := new(lv_dto.IdsReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.IotPrdActionService
	rs := svc.DeleteByIds(req.Ids)
	util.SuccessData(c, rs)
}

// 导出
func (w IotPrdActionController) Export(c *gin.Context) {
	req := new(vo.PageIotPrdActionReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.IotPrdActionService
	url, err := svc.ExportAll(req)
	lv_err.HasErrAndPanic(err)
	util.SucessDataMsg(c, url, url)
}
