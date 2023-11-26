// ==========================================================================
// LV自动生成控制器相关代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2023-11-26 16:27:18 +0800 CST
// 生成人：lv
// ==========================================================================
package controller

import (
    "github.com/gin-gonic/gin"
    "lostvip.com/utils/gconv"
    "lostvip.com/utils/response"
    "robvi/app/modules/biz/model"
    "robvi/app/modules/biz/service"
    sysmodel "robvi/app/modules/sys/model"
)
type DpcTaskController struct{}

///////////////////////////////////////////////////////////////////////////////////////
// 页面部分
// 列表分页数据
///////////////////////////////////////////////////////////////////////////////////////
//查询页
func (w DpcTaskController) List(c *gin.Context) {
	response.BuildTpl(c, "modules/biz/task/list.html").WriteTpl()
}
//新增页
func (w DpcTaskController) Add(c *gin.Context) {
	response.BuildTpl(c, "modules/biz/task/add.html").WriteTpl()
}
//修改页
func (w DpcTaskController) Edit(c *gin.Context) {
	id := gconv.Int64(c.Query("id"))
    taskService :=service.DpcTaskService{}
	entity, err := taskService.SelectRecordById(id)
	if err != nil || entity == nil {
		response.ErrorTpl(c).WriteTpl(gin.H{"desc": "数据不存在",})
		return
	}
	response.BuildTpl(c, "modules/biz/task/edit.html").WriteTpl(gin.H{
		"task": entity,
	})
}

///////////////////////////////////////////////////////////////////////////////////////
// api接口
// 列表分页数据
///////////////////////////////////////////////////////////////////////////////////////
func (w DpcTaskController) ListAjax(c *gin.Context) {
	req := new(model.PageDpcTaskReq)
	if err := c.ShouldBind(req); err != nil {//获取参数
		response.ErrorResp(c).SetMsg(err.Error()).Log("task管理", req).WriteJsonExit()
		return
	}
	rows := make([]model.DpcTask, 0)
	taskService :=service.DpcTaskService{}
	result, page, err := taskService.SelectListByPage(req)

	if err == nil && len(result) > 0 {
		rows = result
	}
	response.BuildTable(c, page.Total, rows).WriteJsonExit()
}


//新增页面保存
func (w DpcTaskController) AddSave(c *gin.Context) {
	req := new(model.AddDpcTaskReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).Log("计划任务新增数据", req).WriteJsonExit()
		return
	}
    taskService :=service.DpcTaskService{}
	id, err := taskService.AddSave(req, c)

	if err != nil || id <= 0 {
		response.ErrorResp(c).Log("计划任务新增数据", req).WriteJsonExit()
		return
	}
	response.SucessResp(c).SetData(id).Log("计划任务新增数据", req).WriteJsonExit()
}


//修改页面保存
func (w DpcTaskController) EditSave(c *gin.Context) {
	req := new(model.EditDpcTaskReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).Log("计划任务修改数据", req).WriteJsonExit()
		return
	}
    taskService :=service.DpcTaskService{}
	rs, err := taskService.EditSave(req, c)

	if err != nil || rs <= 0 {
		response.ErrorResp(c).Log("计划任务修改数据", req).WriteJsonExit()
		return
	}
	response.SucessResp(c).Log("计划任务修改数据", req).WriteJsonExit()
}

//删除数据
func (w DpcTaskController) Remove(c *gin.Context) {
	req := new(sysmodel.RemoveReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).Log("计划任务删除数据", req).WriteJsonExit()
		return
	}
    taskService :=service.DpcTaskService{}
	rs := taskService.DeleteRecordByIds(req.Ids)

	if rs > 0 {
		response.SucessResp(c).Log("计划任务删除数据", req).WriteJsonExit()
	} else {
		response.ErrorResp(c).Log("计划任务删除数据", req).WriteJsonExit()
	}
}

//导出
func (w DpcTaskController) Export(c *gin.Context) {
	req := new(model.PageDpcTaskReq)
	//获取参数
	if err := c.ShouldBind(req); err != nil {
		response.ErrorResp(c).Log("计划任务导出数据", req).WriteJsonExit()
		return
	}

    taskService :=service.DpcTaskService{}
	url, err := taskService.Export(req)

	if err != nil {
		response.ErrorResp(c).Log("计划任务导出数据", req).WriteJsonExit()
		return
	}
	response.SucessResp(c).SetMsg(url).WriteJsonExit()
}