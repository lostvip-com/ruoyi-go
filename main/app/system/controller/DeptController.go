package controller

import (
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_web"
	"net/http"
	"robvi/app/common/session"
	"robvi/app/system/model"
	"robvi/app/system/model/system/dept"
	"robvi/app/system/service"
)

type DeptController struct {
}

// 列表页
func (w *DeptController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "system/dept/list").WriteTpl()
}

// 列表分页数据
func (w *DeptController) ListAjax(c *gin.Context) {
	var service service.DeptService
	var req = dept.SelectPageReq{}
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetMsg(err.Error()).Log("部门管理", req).WriteJsonExit()
		return
	}
	profile := session.GetProfile(c)
	req.TenantId = profile.TenantId

	rows := make([]dept.SysDept, 0)
	result, err := service.SelectListAll(&req)

	if err == nil && len(result) > 0 {
		rows = result
	}

	c.JSON(http.StatusOK, rows)
}

// 新增页面
func (w *DeptController) Add(c *gin.Context) {
	pid := lv_conv.Int64(c.Query("pid"))

	if pid == 0 {
		pid = 100
	}
	service := service.DeptService{}
	tmp := service.SelectDeptById(pid)

	lv_web.BuildTpl(c, "system/dept/add").WriteTpl(gin.H{"dept": tmp})
}

// 新增页面保存
func (w *DeptController) AddSave(c *gin.Context) {
	var req *dept.AddReq
	var service service.DeptService
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg(err.Error()).Log("部门管理", req).WriteJsonExit()
		return
	}

	if service.CheckDeptNameUniqueAll(req.DeptName, req.ParentId) == "1" {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).SetMsg("部门名称已存在").Log("部门管理", req).WriteJsonExit()
		return
	}
	user := session.GetProfile(c)
	isAdmin := session.IsAdminUser(user)
	if isAdmin == false { //非管理员，以当前用户所属租户为限
		req.TenantId = user.TenantId
	}
	rid, err := service.AddSave(req, c)

	if err != nil || rid <= 0 {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Add).Log("部门管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetBtype(model.Buniss_Add).Log("部门管理", req).WriteJsonExit()
}

// 修改页面
func (w *DeptController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	if id <= 0 {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	service := service.DeptService{}
	dept := service.SelectDeptById(id)

	if dept == nil || dept.DeptId <= 0 {
		lv_web.BuildTpl(c, model.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "部门不存在",
		})
		return
	}

	lv_web.BuildTpl(c, "system/dept/edit").WriteTpl(gin.H{
		"dept": dept,
	})
}

// 修改页面保存
func (w *DeptController) EditSave(c *gin.Context) {
	var service service.DeptService

	var req *dept.EditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg(err.Error()).Log("部门管理", req).WriteJsonExit()
		return
	}

	if service.CheckDeptNameUnique(req.DeptName, req.DeptId, req.ParentId) == "1" {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).SetMsg("部门名称已存在").Log("部门管理", req).WriteJsonExit()
		return
	}

	rs, err := service.EditSave(req, c)

	if err != nil || rs <= 0 {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Edit).Log("部门管理", req).WriteJsonExit()
		return
	}
	lv_web.SucessResp(c).SetData(rs).SetBtype(model.Buniss_Edit).Log("部门管理", req).WriteJsonExit()
}

// 删除数据
func (w *DeptController) Remove(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	service := service.DeptService{}
	rs := service.DeleteDeptById(id)

	if rs > 0 {
		lv_web.SucessResp(c).SetBtype(model.Buniss_Del).Log("部门管理", gin.H{"id": id}).WriteJsonExit()
	} else {
		lv_web.ErrorResp(c).SetBtype(model.Buniss_Del).Log("部门管理", gin.H{"id": id}).WriteJsonExit()
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

// 检查部门名称是否已经存在
func (w *DeptController) CheckDeptNameUnique(c *gin.Context) {
	var service service.DeptService
	var req *dept.CheckDeptNameReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}

	result := service.CheckDeptNameUnique(req.DeptName, req.DeptId, req.ParentId)

	c.Writer.WriteString(result)
}

// 检查部门名称是否已经存在
func (w *DeptController) CheckDeptNameUniqueAll(c *gin.Context) {
	var req *dept.CheckDeptNameALLReq
	if err := c.ShouldBind(&req); err != nil {
		c.Writer.WriteString("1")
		return
	}
	var service service.DeptService
	result := service.CheckDeptNameUniqueAll(req.DeptName, req.ParentId)

	c.Writer.WriteString(result)
}