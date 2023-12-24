//	==========================================================================
//
// LV自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：{{.table.CreateTime}}
// 生成人：{{.table.FunctionAuthor}}
// ==========================================================================
package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_logic"
	"lostvip.com/utils/lv_web"
	"robvi/app/common/model_cmn"
	"robvi/app/mywork/service"
	"robvi/app/mywork/vo"
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
	taskService := service.DpcTaskService{}
	entity, err := taskService.FindById(id)
	if err != nil || entity == nil {
		lv_web.ErrorTpl(c).WriteTpl(gin.H{"desc": "数据不存在"})
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
	taskService := service.DpcTaskService{}
	result, total, _ := taskService.ListByPage(req)
	lv_web.SucessPage(c, result, total)
}

// AddSave 新增页面保存
func (w DpcTaskController) AddSave(c *gin.Context) {
	req := new(vo.AddDpcTaskReq)
	err := c.ShouldBind(req)
	lv_logic.HasErrAndPanic(err)
	var userService sysService.UserService
	user := userService.GetProfile(c)
	var svc service.DpcTaskService
	id, err := svc.AddSave(req, user)
	lv_logic.HasErrAndPanic(err)
	lv_web.SucessData(c, id)
}

// EditSave 修改页面保存
func (w DpcTaskController) EditSave(c *gin.Context) {
	req := new(vo.EditDpcTaskReq)
	err := c.ShouldBind(req)
	lv_logic.HasErrAndPanic(err)
	var userService sysService.UserService
	user := userService.GetProfile(c)
	taskService := service.DpcTaskService{}
	err = taskService.EditSave(req, user)
	lv_logic.HasErrAndPanic(err)
	lv_web.Success(c, nil, "success")
}

// Remove 删除数据
func (w DpcTaskController) Remove(c *gin.Context) {
	req := new(model_cmn.RemoveReq)
	if err := c.ShouldBind(req); err != nil {
		panic(err)
		return
	}
	taskService := service.DpcTaskService{}
	rs := taskService.DeleteByIds(req.Ids)
	lv_web.SuccessData(c, rs)
}

// Export 导出
func (w DpcTaskController) Export(c *gin.Context) {
	req := new(vo.PageDpcTaskReq)
	err := c.ShouldBind(req)
	lv_logic.HasErrAndPanic(err)
	taskService := service.DpcTaskService{}
	url, err := taskService.ExportAll(req)
	lv_logic.HasErrAndPanic(err)
	lv_web.SucessDataMsg(c, url, url)
}
