package controller

import (
	"common/cm_vo"
	"common/session"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_web"
	"github.com/lostvip-com/lv_framework/web/dto"
	"main/internal/system/service"
	"net/http"
)

type DeptController struct {
}

// List 列表页
func (w *DeptController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "system/dept/list").WriteTpl()
}

// ListAjax 列表分页数据
func (w *DeptController) ListAjax(c *gin.Context) {
	var service service.DeptService
	var req = cm_vo.DeptPageReq{}
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("部门管理", req).WriteJsonExit()
		return
	}
	result, err := service.SelectListAll(&req)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, result)
}

// Add 新增页面
func (w *DeptController) Add(c *gin.Context) {
	pid := lv_conv.Int64(c.Query("pid"))

	if pid == 0 {
		pid = 100
	}
	service := service.DeptService{}
	tmp := service.SelectDeptById(pid)

	lv_web.BuildTpl(c, "system/dept/add").WriteTpl(gin.H{"dept": tmp})
}

// AddSave 新增页面保存
func (w *DeptController) AddSave(c *gin.Context) {
	var req *cm_vo.AddDeptReq
	var service service.DeptService
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	rid, err := service.AddSave(req, c)
	if err != nil || rid <= 0 {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Add).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetBtype(dto.Buniss_Add).WriteJsonExit()
}

// Edit 修改页面
func (w *DeptController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	service := service.DeptService{}
	dept := service.SelectDeptById(id)

	if dept == nil || dept.DeptId <= 0 {
		lv_web.BuildTpl(c, dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "部门不存在",
		})
		return
	}

	lv_web.BuildTpl(c, "system/dept/edit").WriteTpl(gin.H{
		"dept": dept,
	})
}

// EditSave 修改页面保存
func (w *DeptController) EditSave(c *gin.Context) {
	var service service.DeptService

	var req *cm_vo.EditDeptReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Edit).SetMsg(err.Error()).Log("部门管理", req).WriteJsonExit()
		return
	}

	rs, err := service.EditSave(req, c)

	if err != nil || rs <= 0 {
		lv_web.ErrorResp(c).SetBtype(dto.Buniss_Edit).Log("部门管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetData(rs).SetBtype(dto.Buniss_Edit).Log("部门管理", req).WriteJsonExit()
}

// 删除数据
func (w *DeptController) Remove(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	service := service.DeptService{}
	err := service.DeleteDeptById(id)
	if err != nil {
		lv_web.Fail(c, err.Error())
	} else {
		lv_web.Success(c, id, "success")
	}
}

// 加载部门列表树结构的数据
func (w *DeptController) TreeData(c *gin.Context) {
	var service service.DeptService
	tenantId := session.GetTenantId(c)
	result, _ := service.SelectDeptTree(0, "", "", tenantId)
	c.JSON(http.StatusOK, result)
}

// 加载部门列表树选择页面
func (w *DeptController) SelectDeptTree(c *gin.Context) {
	deptId := lv_conv.Int64(c.Query("deptId"))
	service := service.DeptService{}
	deptPoint := service.SelectDeptById(deptId)

	if deptPoint != nil {
		lv_web.BuildTpl(c, "system/dept/tree").WriteTpl(gin.H{
			"dept": *deptPoint,
		})
	} else {
		lv_web.BuildTpl(c, "system/dept/tree").WriteTpl()
	}
}

// 加载角色部门（数据权限）列表树
func (w *DeptController) RoleDeptTreeData(c *gin.Context) {
	var service service.DeptService
	tenantId := session.GetTenantId(c)
	roleId := lv_conv.Int64(c.Query("roleId"))
	result, err := service.RoleDeptTreeData(roleId, tenantId)

	if err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("菜单树", gin.H{"roleId": roleId})
		return
	}

	c.JSON(http.StatusOK, result)
}
