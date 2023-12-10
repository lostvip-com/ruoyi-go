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
}
