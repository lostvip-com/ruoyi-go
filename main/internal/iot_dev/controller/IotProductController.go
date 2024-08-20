//	=========================================================================
//
// LV自动生成控制器相关代码，只生成一次，按需修改,默认再次生成不会覆盖.
// date：2024-07-19 17:16:13 +0800 CST
// author：lv
// ==========================================================================
package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"main/internal/iot_dev/service"
)

// PrePrdDetailTabs List 查询页
func (w IotProductController) PrePrdDetailTabs(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	productService := service.IotProductService{}
	entity, err := productService.FindById(id)
	if err != nil || entity == nil {
		util.ErrorTpl(c).WriteTpl(gin.H{"desc": "数据不存在"})
		return
	}
	util.WriteTpl(c, "iot_dev/product/prd_detail_tabs.html", gin.H{
		"product": entity,
	})
}

// PreList List 查询页
func (w IotProductController) PreList(c *gin.Context) {
	util.BuildTpl(c, "iot_dev/product/list.html").WriteTpl()
}

// PreAdd 新增页
func (w IotProductController) PreAdd(c *gin.Context) {
	util.BuildTpl(c, "iot_dev/product/add.html").WriteTpl()
}

// PreEdit 修改页
func (w IotProductController) PreEdit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	productService := service.IotProductService{}
	entity, err := productService.FindById(id)
	if err != nil || entity == nil {
		util.ErrorTpl(c).WriteTpl(gin.H{"desc": "数据不存在"})
		return
	}
	util.BuildTpl(c, "iot_dev/product/edit.html").WriteTpl(gin.H{
		"product": entity,
	})
}
