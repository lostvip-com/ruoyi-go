//	=========================================================================
//
// LV自动生成控制器相关代码，只生成一次，按需修改,默认再次生成不会覆盖.
// date：2024-07-19 17:16:13 +0800 CST
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

type IotProductController struct{}

//	========================================================================
//
// api
// =========================================================================

// ListAjax 新增页面保存
func (w IotProductController) ListAjax(c *gin.Context) {
	req := new(vo.PageIotProductReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	if req.PageSize == 0 {
		req.PageSize = lv_dto.PageSize
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	var svc service.IotProductService
	result, total, _ := svc.ListByPage(req)
	util.SucessPage(c, result, total)
}

// AddSave 新增页面保存
func (w IotProductController) AddSave(c *gin.Context) {
	req := new(vo.AddIotProductReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.IotProductService

	var userService sysService.UserService
	user := userService.GetProfile(c)
	req.CreateBy = user.LoginName
	id, err := svc.AddSave(req)
	lv_err.HasErrAndPanic(err)
	util.SucessData(c, id)
}

// EditSave 修改页面保存
func (w IotProductController) EditSave(c *gin.Context) {
	req := new(vo.EditIotProductReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.IotProductService
	var userService sysService.UserService
	user := userService.GetProfile(c)
	req.UpdateBy = user.LoginName
	err = svc.EditSave(req)
	lv_err.HasErrAndPanic(err)
	util.Success(c, nil, "success")
}

// Remove 删除数据
func (w IotProductController) Remove(c *gin.Context) {
	req := new(lv_dto.IdsReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.IotProductService
	rs := svc.DeleteByIds(req.Ids)
	util.SuccessData(c, rs)
}

// 导出
func (w IotProductController) Export(c *gin.Context) {
	req := new(vo.PageIotProductReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.IotProductService
	url, err := svc.ExportAll(req)
	lv_err.HasErrAndPanic(err)
	util.SucessDataMsg(c, url, url)
}
