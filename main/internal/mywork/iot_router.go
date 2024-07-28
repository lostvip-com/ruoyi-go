package mywork

import (
	"fmt"
	"github.com/lostvip-com/lv_framework/web/router"
	"main/internal/mywork/controller"
)

func init() {
	fmt.Println("############## iot_dev init ########################")
	//g1 := router.New( "/demo/form",token.JWTAuthMiddleware())
	hisData := controller.HisDataController{}
	//g1 := router.New("/iot_dev/hisData")
	g1 := router.New("/")
	//lv_web
	g1.GET("/iot_dev/hisData/toWizard", "", hisData.ToWizard)
	//data
	g1.GET("/iot_dev/hisData/sugSite", "", hisData.SugSite)
	//反向代理
	g1.POST("/datasyn/csv2HisController/copyCsv2His", "", hisData.CopyCsv2His)
	g1.POST("/datasyn/hisDataController/copy_his_2_aggDevPowerDay", "", hisData.CopyHis2AggDevPowerDay)
	g1.POST("/datasyn/hisDataController/copy_his_2_UnqSeqPower", "", hisData.CopyHis2UnqSeqPower)

}
