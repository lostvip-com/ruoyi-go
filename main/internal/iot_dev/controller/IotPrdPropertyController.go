//	=========================================================================
//
// LV自动生成控制器相关代码，只生成一次，按需修改,默认再次生成不会覆盖.
// date：2024-07-30 21:59:38 +0800 CST
// author：lv
// ==========================================================================
package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"main/internal/iot_dev/service"
)

// List 查询页
func (w IotPrdPropertyController) List(c *gin.Context) {
	productId := c.Query("productId")
	util.BuildTpl(c, "iot_dev/property/list.html").WriteTpl(gin.H{
		"productId": productId,
	})
}

// Add 新增页
func (w IotPrdPropertyController) Add(c *gin.Context) {
	util.BuildTpl(c, "iot_dev/property/add.html").WriteTpl()
}

// Edit 修改页
func (w IotPrdPropertyController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	propertyService := service.IotPrdPropertyService{}
	entity, err := propertyService.FindById(id)
	if err != nil || entity == nil {
		util.ErrorTpl(c).WriteTpl(gin.H{"desc": "数据不存在"})
		return
	}
	util.BuildTpl(c, "iot_dev/property/edit.html").WriteTpl(gin.H{
		"property": entity,
	})
}
