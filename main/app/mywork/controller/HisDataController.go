package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lostvip.com/utils/lv_net"
	"net/url"
)

type HisDataController struct {
}

/**
 * 测试使用，代理本地服务
 */
//func (w HisDataController) SugSite(c *gin.Context) {
//	fmt.Println(" ########## host: ")
//	host := "127.0.0.1:7300"
//	lv_net.ProxyWithUrlSame(c, host)
//	fmt.Println("=============over==================")
//}

/**
 * 代理网关地址
 */
func (w HisDataController) SugSite(c *gin.Context) {
	fmt.Println("=============SugSite start==================")
	query := "q=" + url.QueryEscape(c.Query("q"))
	url := "http://iot-gateway:8888/iotsettings/navtree/site"
	lv_net.ProxyWithUrlDifferent(c, url, query)
	fmt.Println("=============SugSite over==================")
}

/**
 * 代理网关地址
 * "http://gateway1:8888/datasyn/hisDataController/copy_his_2_aggDevPowerDay"
 */
func (w HisDataController) CopyCsv2His(c *gin.Context) {
	host := "iot-gateway:8888"
	fmt.Println("=============copy_his_2_aggDevPowerDay start================" + host)
	lv_net.ProxyWithUrlSame(c, host)
	fmt.Println("=========copy_his_2_aggDevPowerDay over==================")
}

/**
 * 代理网关地址
 * "http://gateway1:8888/datasyn/hisDataController/copy_his_2_aggDevPowerDay"
 */
func (w HisDataController) CopyHis2AggDevPowerDay(c *gin.Context) {
	host := "gateway1:8888"
	fmt.Println("=============copy_his_2_aggDevPowerDay start================" + host)
	lv_net.ProxyWithUrlSame(c, host)
	fmt.Println("=========copy_his_2_aggDevPowerDay over==================")
}

/**
 *  重新计算 结算数据
 *
 * 代理网关地址
 * "http://gateway1:8888/datasyn/hisDataController/copy_his_2_UnqSeqPower"
 */
func (w HisDataController) CopyHis2UnqSeqPower(c *gin.Context) {
	host := "gateway1:8888"
	fmt.Println("=============copy_his_2_aggDevPowerDay start================" + host)
	lv_net.ProxyWithUrlSame(c, host)
	fmt.Println("=========copy_his_2_aggDevPowerDay over==================")
}
