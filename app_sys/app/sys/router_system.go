package sys

import (
	"lostvip.com/web/router"
	"robvi/app/common/middleware/auth"
	"robvi/app/common/middleware/token"
	controller2 "robvi/app/sys/controller"
	"robvi/app/sys/controller/system/dict_data"
	"robvi/app/sys/controller/system/dict_type"
)

func init() {
	// 加载框架路由
	g0 := router.New("/system", token.TokenMiddleware(), auth.Auth)
	index := controller2.MainController{}
	g0.GET("/main", "", index.Main)
	g0.GET("/switchSkin", "", index.SwitchSkin)
	g0.GET("/download", "", index.Download)
	//系统配置
	g1 := router.New("/system/config", token.TokenMiddleware(), auth.Auth)
	config := controller2.ConfigController{}
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
	g3.GET("/", "system:dict:view", dict_type.List)
	g3.POST("/list", "system:dict:list", dict_type.ListAjax)
	g3.GET("/add", "system:dict:add", dict_type.Add)
	g3.POST("/add", "system:dict:add", dict_type.AddSave)
	g3.POST("/remove", "system:dict:remove", dict_type.Remove)
	g3.GET("/remove", "system:dict:remove", dict_type.Remove)
	g3.GET("/edit", "system:dict:edit", dict_type.Edit)
	g3.POST("/edit", "system:dict:edit", dict_type.EditSave)
	g3.GET("/detail", "system:dict:detail", dict_type.Detail)
	g3.POST("/export", "system:dict:export", dict_type.Export)
	g3.POST("/checkDictTypeUniqueAll", "system:dict:view", dict_type.CheckDictTypeUniqueAll)
	g3.POST("/checkDictTypeUnique", "system:dict:view", dict_type.CheckDictTypeUnique)
	g3.GET("/selectDictTree", "system:dict:view", dict_type.SelectDictTree)
	g3.GET("/treeData", "system:dict:view", dict_type.TreeData)
	// 字典内容参数路由
	g4 := router.New("/system/dict/data", token.TokenMiddleware(), auth.Auth)
	g4.POST("/list", "system:dict:view", dict_data.ListAjax)
	g4.GET("/add", "system:dict:add", dict_data.Add)
	g4.POST("/add", "system:dict:add", dict_data.AddSave)
	g4.POST("/remove", "system:dict:remove", dict_data.Remove)
	g4.GET("/edit", "system:dict:edit", dict_data.Edit)
	g4.POST("/edit", "system:dict:edit", dict_data.EditSave)
	g4.POST("/export", "system:dict:export", dict_data.Export)
}
