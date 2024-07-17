//	=========================================================================
//
// LV自动生成控制器相关代码，只生成一次，按需修改,默认再次生成不会覆盖.
// date：2024-02-28 14:21:50 +0800 CST
// author：lv
// ==========================================================================
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_web"
	"github.com/lostvip-com/lv_framework/web/dto"
	"github.com/spf13/cast"
	"main/internal/mywork/model"
	"main/internal/mywork/service"
	"main/internal/mywork/vo"
	sysService "main/internal/system/service"
	"strings"
	"time"
)

type AppGenParamsController struct{}

//	========================================================================
//
// api
// =========================================================================

// AddSave 新增页面保存
func (w AppGenParamsController) GenParams(c *gin.Context) {
	baseNum := c.PostForm("baseNum")
	amount := c.PostForm("amount")
	var svc service.AppGenParamsService
	var userService sysService.UserService
	user := userService.GetProfile(c)
	svc.GenParamsToDB(cast.ToInt(baseNum), cast.ToInt(amount), user.LoginName)
	lv_web.SucessData(c, "success")
}

// 修改页面保存
func (w *AppGenParamsController) ChangeProp(c *gin.Context) {
	id := c.PostForm("id")
	value := c.PostForm("value")
	name := c.PostForm("name")
	po := model.AppGenParams{Id: cast.ToInt64(id)}
	err := po.FindById()
	lv_err.HasErrAndPanic(err)
	if po.ParamName != "" && po.UseFlag == "1" && po.MonitorTypeId > 0 {
		if strings.Contains(name, "remark") {
			po.Remark = value
			lv_err.HasErrAndPanic(err)
			lv_web.SuccessData(c, po)
		} else {
			panic("已经绑定业务，不允许修改")
		}
	} else {
		if strings.Contains(name, "monitorTypeId") {
			po.MonitorTypeId = cast.ToInt(value)
		} else if strings.Contains(name, "paramName") {
			po.ParamName = value
		} else if strings.Contains(name, "remark") {
			po.Remark = value
		} else if strings.Contains(name, "unit") {
			po.Unit = value
		} else if strings.Contains(name, "paramType") {
			po.ParamType = value
		} else {
			panic("已经绑定业务，不允许修改")
		}
		po.UpdateTime = time.Now()
		po.UseFlag = "0"
		err = po.Save()
		lv_err.HasErrAndPanic(err)
		lv_web.SuccessData(c, po)
	}

}

// 修改页面保存
func (w *AppGenParamsController) ChangeStatus(c *gin.Context) {
	id := c.PostForm("id")
	status := c.PostForm("status")
	po := model.AppGenParams{Id: cast.ToInt64(id)}
	err := po.FindById()
	lv_err.HasErrAndPanic(err)
	if po.ParamName != "" && po.UseFlag == "1" && po.MonitorTypeId > 0 {
		panic("已经绑定业务，不允许修改状态")
	}
	if po.ParamName == "" && po.MonitorTypeId == 0 {
		panic("未设置监控类型和业务名称，不允许修改状态")
	}
	po.UpdateTime = time.Now()
	po.UseFlag = status
	err = po.Save()
	lv_err.HasErrAndPanic(err)
	lv_web.SuccessData(c, po)
}

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
	req := new(dto.IdsReq)
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
	lv_web.SucessDataMsg(c, url, "")
}
