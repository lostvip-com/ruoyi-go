package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	sysService "system/service"
	"{{.ModuleName}}/internal/{{.PackageName}}/service"
	"{{.ModuleName}}/internal/{{.PackageName}}/vo"
)

type {{.ClassName}}Api struct {
}

//	========================================================================
//
// api
// =========================================================================

// List{{.ClassName}} 新增页面保存
func (w {{.ClassName}}Api) List{{.ClassName}}(c *gin.Context) {
	req := new(vo.{{.ClassName}}Req)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	//if req.DeptId == 0 {
	//	req.DeptId = session.GetLoginInfo(c).DeptId
	//}
	var svc service.{{.ClassName}}Service
	result, total, _ := svc.ListByPage(req)
	util.SucessPage(c, result, total)
}

// Create{{.ClassName}} 新增页面保存
func (w {{.ClassName}}Api) Add{{.ClassName}}(c *gin.Context) {
	req := new(vo.Add{{.ClassName}}Req)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.{{.ClassName}}Service
	var userService sysService.UserService
	user := userService.GetProfile(c)
	req.CreateBy = user.LoginName
	id, err := svc.AddSave(req)
	lv_err.HasErrAndPanic(err)
	util.SucessData(c, id)
}

// Save{{.ClassName}} 修改页面保存
func (w {{.ClassName}}Api) Save{{.ClassName}}(c *gin.Context) {
	req := new(vo.Edit{{.ClassName}}Req)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.{{.ClassName}}Service
	var userService sysService.UserService
	user := userService.GetProfile(c)
	req.UpdateBy = user.LoginName
	err = svc.EditSave(req)
	lv_err.HasErrAndPanic(err)
	util.Success(c, nil, "success")
}

// Remove{{.ClassName}} 删除数据
func (w {{.ClassName}}Api) Remove{{.ClassName}}(c *gin.Context) {
	req := new(lv_dto.IdsReq)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.{{.ClassName}}Service
	rs := svc.DeleteByIds(req.Ids)
	util.SuccessData(c, rs)
}

// 导出
func (w {{.ClassName}}Api) Export{{.ClassName}}(c *gin.Context) {
	req := new(vo.{{.ClassName}}Req)
	err := c.ShouldBind(req)
	lv_err.HasErrAndPanic(err)
	var svc service.{{.ClassName}}Service
	url, err := svc.ExportAll(req)
	lv_err.HasErrAndPanic(err)
	util.SucessDataMsg(c, url, url)
}

