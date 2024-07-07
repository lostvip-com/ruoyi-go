//	=========================================================================
//
// LV自动生成控制器相关代码，只生成一次，按需修改,默认再次生成不会覆盖.
// date：2024-02-28 14:21:50 +0800 CST
// author：lv
// ==========================================================================
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lv_framework/utils/lv_conv"
	"github.com/lv_framework/utils/lv_web"
	"main/app/mywork/service"
)

// List 查询页
func (w AppGenParamsController) List(c *gin.Context) {
	lv_web.BuildTpl(c, "mywork/params/list.html").WriteTpl()
}

// Add 新增页
func (w AppGenParamsController) Add(c *gin.Context) {
	lv_web.BuildTpl(c, "mywork/params/add.html").WriteTpl()
}

// Edit 修改页
func (w AppGenParamsController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	paramsService := service.AppGenParamsService{}
	entity, err := paramsService.FindById(id)
	if err != nil || entity == nil {
		lv_web.ErrorTpl(c).WriteTpl(gin.H{"desc": "数据不存在"})
		return
	}
	lv_web.BuildTpl(c, "mywork/params/edit.html").WriteTpl(gin.H{
		"params": entity,
	})
}
