package controller

import (
	"common/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_net"
	"net/url"
)

type IndexController struct {
}

func (w IndexController) Index(c *gin.Context) {
	util.BuildTpl(c, "index.html").WriteTpl()
}

/**
 *  更换请求链接
 *  http://iot-gateway:8888/iotsettings/navtree/site
 */

func (w IndexController) SugSite(c *gin.Context) {
	fmt.Println("=============SugSite start==================")
	query := "q=" + url.QueryEscape(c.Query("q"))
	url := "http://iot-gateway:8888/iotsettings/navtree/site"
	lv_net.ProxyWithUrlDifferent(c, url, query)
	fmt.Println("=============SugSite over==================")
}

/**
 * 只更换代理网关 ip和port
 * "http://gateway1:8888/datasyn/hisDataController/copy_his_2_aggDevPowerDay"
 */

func (w IndexController) CopyCsv2His(c *gin.Context) {
	host := "iot_dev-gateway:8888"
	fmt.Println("=============copy_his_2_aggDevPowerDay start================" + host)
	lv_net.ProxyWithUrlSame(c, host)
	fmt.Println("=========copy_his_2_aggDevPowerDay over==================")
}

/**
 *   代理地址
 *  "http://gateway1:8888/datasyn/hisDataController/copy_his_2_aggDevPowerDay"
 */

func (w IndexController) CopyHis2AggDevPowerDay(c *gin.Context) {
	host := "gateway1:8888"
	fmt.Println("=============copy_his_2_aggDevPowerDay start================" + host)
	lv_net.ProxyWithUrlSame(c, host)
	fmt.Println("=========copy_his_2_aggDevPowerDay over==================")
}
