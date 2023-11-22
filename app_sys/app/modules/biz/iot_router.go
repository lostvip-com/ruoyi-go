package robvi

import (
	"fmt"
	"lostvip.com/web/router"
	"robvi/app/modules/biz/controller"
)

func init() {
	fmt.Println("############## iot init ########################")
	//g1 := router.New( "/demo/form",token.JWTAuthMiddleware())
	hisData := controller.HisDataController{}
	//g1 := router.New("/iot/hisData")
	g1 := router.New("/")
	//page
	g1.GET("/iot/hisData/toWizard", "", hisData.ToWizard)
	//data
	g1.GET("/iot/hisData/sugSite", "", hisData.SugSite)
	//反向代理
	g1.POST("/datasyn/csv2HisController/copyCsv2His", "", hisData.CopyCsv2His)
	g1.POST("/datasyn/hisDataController/copy_his_2_aggDevPowerDay", "", hisData.CopyHis2AggDevPowerDay)
	g1.POST("/datasyn/hisDataController/copy_his_2_UnqSeqPower", "", hisData.CopyHis2UnqSeqPower)

}
