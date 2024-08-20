//	=========================================================================
//
// LV自动生成控制器相关代码，只生成一次，按需修改,默认再次生成不会覆盖.
// date：2024-02-28 14:21:50 +0800 CST
// author：lv
// ==========================================================================
package controller

import (
	util "common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"main/internal/mywork/service"
)

// List 查询页
func (w AppGenParamsController) List(c *gin.Context) {
	util.BuildTpl(c, "mywork/params/list.html").WriteTpl()
}

// Add 新增页
func (w AppGenParamsController) Add(c *gin.Context) {
	util.WriteTpl(c, "mywork/params/add.html")
}

// Edit 修改页
func (w AppGenParamsController) Edit(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
	paramsService := service.AppGenParamsService{}
	entity, err := paramsService.FindById(id)
	if err != nil || entity == nil {
		util.ErrorTpl(c).WriteTpl(gin.H{"desc": "数据不存在"})
		return
	}
	util.WriteTpl(c, "mywork/params/edit.html", gin.H{"params": entity})
}
