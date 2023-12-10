package demo

import (
	"fmt"
	"lostvip.com/web/router"
	"robvi/app/common/middleware/token"
	controller2 "robvi/app/demo/controller"
)

func init() {
	fmt.Println("############## demo init ################")
	//g1 := router.New( "/demo/form",token.TokenMiddleware())
	demo := controller2.DemoController{}
	g0 := router.New("/demo/db")
	//mybatis
	g0.GET("/mybatis1", "", demo.Mybatis1)
	g0.GET("/mybatis2", "", demo.Mybatis2)
	g0.GET("/mybatis3", "", demo.Mybatis3)
	g0.GET("/redis", "", demo.TestRedis)
	//lv_web
	g1 := router.New("/demo/form")
	g1.GET("/autocomplete", "", demo.Autocomplete)
	g1.GET("/basic", "", demo.Basic)
	g1.GET("/button", "", demo.Button)
	g1.GET("/cards", "", demo.Cards)
	g1.GET("/datetime", "", demo.Datetime)
	g1.GET("/duallistbox", "", demo.Duallistbox)
	g1.GET("/grid", "", demo.Grid)

	g1.GET("/jasny", "", demo.Jasny)
	g1.GET("/select", "", demo.Select)
	g1.GET("/sortable", "", demo.Sortable)
	g1.GET("/summernote", "", demo.Summernote)
	g1.GET("/tabs_panels", "", demo.Tabs_panels)

	g1.GET("/timeline", "", demo.Timeline)
	g1.GET("/upload", "", demo.Upload)
	g1.GET("/validate", "", demo.Validate)
	g1.GET("/wizard", "", demo.Wizard)

	g2 := router.New("/demo/icon", token.TokenMiddleware())
	g2.GET("/fontawesome", "", demo.Fontawesome)
	g2.GET("/glyphicons", "", demo.Glyphicons)

	g3 := router.New("/demo/modal", token.TokenMiddleware())
	g3.GET("/dialog", "", demo.Dialog)
	g3.GET("/form", "", demo.Form)
	g3.GET("/layer", "", demo.Layer)
	g3.GET("/table", "", demo.Table)
	g3.GET("/check", "", demo.Check)
	g3.GET("/parent", "", demo.Parent)
	g3.GET("/radio", "", demo.Radio)

	g4 := router.New("/demo/operate", token.TokenMiddleware())
	oper := controller2.OperateController{}
	g4.GET("/list", "", oper.List)
	g4.GET("/add", "", oper.Add)
	g4.GET("/detail", "", oper.Detail)
	g4.GET("/edit", "", oper.Edit)
	g4.POST("/edit", "", oper.EditSave)
	g4.GET("/other", "", oper.Other)
	g4.GET("/table", "", oper.Table)

	g5 := router.New("/demo/report", token.TokenMiddleware())
	g5.GET("/echarts", "", demo.Echarts)
	g5.GET("/metrics", "", demo.Metrics)
	g5.GET("/peity", "", demo.Peity)
	g5.GET("/sparkline", "", demo.Sparkline)

	g6 := router.New("/demo/table", token.TokenMiddleware())
	g6.GET("/button", "", demo.Button)
	g6.GET("/child", "", demo.Child)
	g6.GET("/curd", "", demo.Curd)
	g6.GET("/detail", "", demo.Detail)
	g6.POST("list", "", demo.List)

	g6.GET("/editable", "", demo.Editable, token.TokenMiddleware())
	g6.GET("/event", "", demo.Event)
	g6.POST("/export", "", demo.Export)
	g6.GET("/fixedColumns", "", demo.FixedColumns)
	g6.GET("/footer", "", demo.Footer)
	g6.GET("/groupHeader", "", demo.GroupHeader)

	g6.GET("/image", "", demo.Image)
	g6.GET("/multi", "", demo.Multi)
	g6.GET("/other", "", demo.Other)
	g6.GET("/pageGo", "", demo.PageGo)

	g6.GET("/params", "", demo.Params)
	g6.GET("/remember", "", demo.Remember)
	g6.GET("/recorder", "", demo.Recorder)
	g6.GET("/lv_sql", "", demo.Search)
}
