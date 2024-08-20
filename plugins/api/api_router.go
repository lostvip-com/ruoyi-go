package api

import (
	"fmt"
	"github.com/lostvip-com/lv_framework/web/router"
	"plugins/api/internal/mywork/controller"
)

func init() {
	fmt.Println("############## plugins api init ########################")
	apiIndex := controller.IndexController{}
	g1 := router.New("/plugin/api")
	// page
	g1.GET("/", "", apiIndex.Index)
	// api
	g1.GET("/sugSite", "", apiIndex.SugSite)

	//反向代理
	g1.POST("/CopyCsv2His", "", apiIndex.CopyCsv2His)
	g1.POST("/CopyHis2AggDevPowerDay", "", apiIndex.CopyHis2AggDevPowerDay)

}
