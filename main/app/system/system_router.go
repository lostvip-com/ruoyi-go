package system

import (
	"lostvip.com/web/router"
	"robvi/app/common/middleware/auth"
	"robvi/app/common/middleware/token"
	controller "robvi/app/system/controller"
)

func init() {
	// 加载框架路由
	g0 := router.New("/system", token.TokenMiddleware(), auth.Auth)
	index := controller.MainController{}
	g0.GET("/main", "", index.Main)
	g0.GET("/switchSkin", "", index.SwitchSkin)
	g0.GET("/download", "", index.Download)
	//系统配置
	g1 := router.New("/system/config", token.TokenMiddleware(), auth.Auth)
	config := controller.ConfigController{}
	g1.GET("/", "system:config:view", config.List)
	g1.POST("/list", "system:config:list", config.ListAjax)
	g1.GET("/add", "system:config:add", config.Add)
	g1.POST("/add", "system:config:add", config.AddSave)
	g1.POST("/remove", "system:config:remove", config.Remove)
	g1.GET("/edit", "system:config:edit", config.Edit)
	g1.POST("/edit", "system:config:edit", config.EditSave)
	g1.POST("/export", "system:config:export", config.Export)
	g1.POST("/checkConfigKeyUniqueAll", "system:config:view", config.CheckConfigKeyUniqueAll)
	g1.POST("/checkConfigKeyUnique", "system:config:view", config.CheckConfigKeyUnique)
	// 字典类型参数路由
	g3 := router.New("/system/dict", token.TokenMiddleware(), auth.Auth)
	dictType := controller.DictTypeController{}
	g3.GET("/", "system:dict:view", dictType.List)
	g3.POST("/list", "system:dict:list", dictType.ListAjax)
	g3.GET("/add", "system:dict:add", dictType.Add)
	g3.POST("/add", "system:dict:add", dictType.AddSave)
	g3.POST("/remove", "system:dict:remove", dictType.Remove)
	g3.GET("/remove", "system:dict:remove", dictType.Remove)
	g3.GET("/edit", "system:dict:edit", dictType.Edit)
	g3.POST("/edit", "system:dict:edit", dictType.EditSave)
	g3.GET("/detail", "system:dict:detail", dictType.Detail)
	g3.POST("/export", "system:dict:export", dictType.Export)
	g3.POST("/checkDictTypeUniqueAll", "system:dict:view", dictType.CheckDictTypeUniqueAll)
	g3.POST("/checkDictTypeUnique", "system:dict:view", dictType.CheckDictTypeUnique)
	g3.GET("/selectDictTree", "system:dict:view", dictType.SelectDictTree)
	g3.GET("/treeData", "system:dict:view", dictType.TreeData)
	// 字典内容参数路由
	g4 := router.New("/system/dict/data", token.TokenMiddleware(), auth.Auth)
	dictData := controller.DictDataController{}
	g4.POST("/list", "system:dict:view", dictData.ListAjax)
	g4.GET("/add", "system:dict:add", dictData.Add)
	g4.POST("/add", "system:dict:add", dictData.AddSave)
	g4.POST("/remove", "system:dict:remove", dictData.Remove)
	g4.GET("/edit", "system:dict:edit", dictData.Edit)
	g4.POST("/edit", "system:dict:edit", dictData.EditSave)
	g4.POST("/export", "system:dict:export", dictData.Export)
	//dept
	groupDept := router.New("/system/dept", token.TokenMiddleware(), auth.Auth)
	deptContr := controller.DeptController{}
	groupDept.GET("/", "system:dept:view", deptContr.List)
	groupDept.POST("/list", "system:dept:list", deptContr.ListAjax)
	groupDept.GET("/add", "system:dept:add", deptContr.Add)
	groupDept.POST("/add", "system:dept:add", deptContr.AddSave)
	groupDept.POST("/remove", "system:dept:remove", deptContr.Remove)
	groupDept.GET("/remove", "system:dept:remove", deptContr.Remove)
	groupDept.GET("/edit", "system:dept:edit", deptContr.Edit)
	groupDept.POST("/edit", "system:dept:edit", deptContr.EditSave)
	groupDept.POST("/checkDeptNameUnique", "system:dept:view", deptContr.CheckDeptNameUnique)
	groupDept.POST("/checkDeptNameUniqueAll", "system:dept:view", deptContr.CheckDeptNameUniqueAll)
	groupDept.GET("/treeData", "system:dept:view", deptContr.TreeData)
	groupDept.GET("/selectDeptTree", "system:dept:view", deptContr.SelectDeptTree)
	groupDept.GET("/roleDeptTreeData", "system:dept:view", deptContr.RoleDeptTreeData)
	// 用户管理路由
	groupUser := router.New("/system/user", token.TokenMiddleware(), auth.Auth)
	user := controller.UserController{}
	groupUser.GET("/", "system:user:view", user.List)
	groupUser.POST("/list", "system:user:list", user.ListAjax)
	groupUser.GET("/add", "system:user:add", user.Add)
	groupUser.POST("/add", "system:user:add", user.AddSave)
	groupUser.POST("/remove", "system:user:remove", user.Remove)
	groupUser.GET("/edit", "system:user:edit", user.Edit)
	groupUser.POST("/edit", "system:user:edit", user.EditSave)
	groupUser.POST("/export", "system:user:export", user.Export)
	groupUser.GET("/resetPwd", "system:user:resetPwd", user.ResetPwd)
	groupUser.POST("/resetPwd", "system:user:resetPwd", user.ResetPwdSave)

	// 个人中心路由
	groupProfile := router.New("/system/user/profile", token.TokenMiddleware(), auth.Auth)
	profile := controller.ProfileController{}
	groupProfile.GET("/", "", profile.Profile)
	groupProfile.GET("/avatar", "", profile.Avatar)
	groupProfile.GET("/resetPwd", "", profile.EditPwd)
	groupProfile.POST("/update", "", profile.Update)
	groupProfile.POST("/resetSavePwd", "", profile.UpdatePassword)
	groupProfile.POST("/checkEmailUnique", "", profile.CheckEmailUnique)
	groupProfile.POST("/checkPhoneUnique", "", profile.CheckPhoneUnique)
	groupProfile.POST("/checkLoginNameUnique", "", profile.CheckLoginNameUnique)
	groupProfile.POST("/checkEmailUniqueAll", "", profile.CheckEmailUniqueAll)
	groupProfile.POST("/checkPhoneUniqueAll", "", profile.CheckPhoneUniqueAll)
	groupProfile.POST("/checkPassword", "", profile.CheckPassword)
	groupProfile.POST("/updateAvatar", "", profile.UpdateAvatar)
	// 角色路由
	// 角色路由
	groupRole := router.New("/system/role", token.TokenMiddleware(), auth.Auth)
	roleController := controller.RoleController{}
	groupRole.GET("/", "system:post:view", roleController.List)
	groupRole.POST("/list", "system:post:list", roleController.ListAjax)
	groupRole.GET("/add", "system:post:add", roleController.Add)
	groupRole.POST("/add", "system:post:add", roleController.AddSave)
	groupRole.POST("/remove", "system:post:remove", roleController.Remove)
	groupRole.GET("/edit", "system:post:edit", roleController.Edit)
	groupRole.POST("/edit", "system:post:edit", roleController.EditSave)
	groupRole.POST("/export", "system:post:export", roleController.Export)
	groupRole.POST("/checkRoleKeyUnique", "system:post:view", roleController.CheckRoleKeyUnique)
	groupRole.POST("/checkRoleNameUniqueAll", "system:post:view", roleController.CheckRoleNameUniqueAll)
	groupRole.POST("/checkRoleNameUnique", "system:post:view", roleController.CheckRoleNameUnique)
	groupRole.POST("/checkRoleKeyUniqueAll", "system:post:view", roleController.CheckRoleKeyUniqueAll)
	groupRole.GET("/authDataScope", "system:post:view", roleController.AuthDataScope)
	groupRole.POST("/authDataScope", "system:post:view", roleController.AuthDataScopeSave)
	groupRole.GET("/authUser", "system:post:view", roleController.AuthUser)
	groupRole.POST("/allocatedList", "system:post:view", roleController.AllocatedList)
	groupRole.GET("/selectUser", "system:post:view", roleController.SelectUser)
	groupRole.POST("/unallocatedList", "system:post:view", roleController.UnallocatedList)
	groupRole.POST("/selectAll", "system:post:view", roleController.SelectAll)
	groupRole.POST("/cancel", "system:post:view", roleController.Cancel)
	groupRole.POST("/cancelAll", "system:post:view", roleController.CancelAll)
	// 菜单路由
	groupMenu := router.New("/system/menu", token.TokenMiddleware(), auth.Auth)
	menuController := controller.MenuController{}
	groupMenu.GET("/", "system:menu:view", menuController.List)
	groupMenu.POST("/list", "system:menu:list", menuController.ListAjax)
	groupMenu.GET("/add", "system:menu:add", menuController.Add)
	groupMenu.POST("/add", "system:menu:add", menuController.AddSave)
	groupMenu.GET("/remove", "system:menu:remove", menuController.Remove)
	groupMenu.POST("/remove", "system:menu:remove", menuController.Remove)
	groupMenu.GET("/edit", "system:menu:edit", menuController.Edit)
	groupMenu.POST("/edit", "system:menu:edit", menuController.EditSave)
	groupMenu.GET("/icon", "system:menu:view", menuController.Icon)
	groupMenu.GET("/selectMenuTree", "system:menu:view", menuController.SelectMenuTree)
	groupMenu.GET("/roleMenuTreeData", "system:menu:view", menuController.RoleMenuTreeData)
	groupMenu.GET("/menuTreeData", "system:menu:view", menuController.MenuTreeData)
	groupMenu.POST("/checkMenuNameUnique", "system:menu:view", menuController.CheckMenuNameUnique)
	groupMenu.POST("/checkMenuNameUniqueAll", "system:menu:view", menuController.CheckMenuNameUniqueAll)
	// 岗位路由
	groupPost := router.New("/system/post", token.TokenMiddleware(), auth.Auth)
	postController := controller.PostController{}
	groupPost.GET("/", "system:post:view", postController.List)
	groupPost.POST("/list", "system:post:list", postController.ListAjax)
	groupPost.GET("/add", "system:post:add", postController.Add)
	groupPost.POST("/add", "system:post:add", postController.AddSave)
	groupPost.POST("/remove", "system:post:remove", postController.Remove)
	groupPost.GET("/edit", "system:post:edit", postController.Edit)
	groupPost.POST("/edit", "system:post:edit", postController.EditSave)
	groupPost.POST("/export", "system:post:export", postController.Export)
	groupPost.POST("/checkPostCodeUniqueAll", "system:post:list", postController.CheckPostCodeUniqueAll)
	groupPost.POST("/checkPostCodeUnique", "system:post:list", postController.CheckPostCodeUnique)
	groupPost.POST("/checkPostNameUniqueAll", "system:post:list", postController.CheckPostNameUniqueAll)
	groupPost.POST("/checkPostNameUnique", "system:post:list", postController.CheckPostNameUnique)
}
