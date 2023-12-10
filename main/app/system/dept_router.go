package system

import (
	"lostvip.com/web/router"
	"robvi/app/common/middleware/auth"
	"robvi/app/common/middleware/token"
	"robvi/app/system/controller"
)

// 加载路由
func init() {
	// 分组路由注册方式
	g1 := router.New("/system/dept", token.TokenMiddleware(), auth.Auth)
	deptContr := controller.DeptController{}
	g1.GET("/", "system:dept:view", deptContr.List)
	g1.POST("/list", "system:dept:list", deptContr.ListAjax)
	g1.GET("/add", "system:dept:add", deptContr.Add)
	g1.POST("/add", "system:dept:add", deptContr.AddSave)
	g1.POST("/remove", "system:dept:remove", deptContr.Remove)
	g1.GET("/remove", "system:dept:remove", deptContr.Remove)
	g1.GET("/edit", "system:dept:edit", deptContr.Edit)
	g1.POST("/edit", "system:dept:edit", deptContr.EditSave)
	g1.POST("/checkDeptNameUnique", "system:dept:view", deptContr.CheckDeptNameUnique)
	g1.POST("/checkDeptNameUniqueAll", "system:dept:view", deptContr.CheckDeptNameUniqueAll)
	g1.GET("/treeData", "system:dept:view", deptContr.TreeData)
	g1.GET("/selectDeptTree", "system:dept:view", deptContr.SelectDeptTree)
	g1.GET("/roleDeptTreeData", "system:dept:view", deptContr.RoleDeptTreeData)
}
