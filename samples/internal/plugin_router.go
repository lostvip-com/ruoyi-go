package internal

import (
	"fmt"
	"github.com/lostvip-com/lv_framework/web/router"
	"simples/internal/mywork/controller"
)

func init() {
	fmt.Println("############## init ########################")
	hisData := controller.IndexController{}
	g1 := router.New("/")
	// page
	g1.GET("/", "", hisData.Index)

	// api
	g1.GET("/sugSite", "", hisData.SugSite)

	//反向代理
	g1.POST("/CopyCsv2His", "", hisData.CopyCsv2His)
	g1.POST("/CopyHis2AggDevPowerDay", "", hisData.CopyHis2AggDevPowerDay)

}
