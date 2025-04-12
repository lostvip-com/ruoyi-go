//	=========================================================================
//
// LV自动生成控制器相关代码,只生成一次,按需修改,默认再次生成不会覆盖.
// date: {{.CreateTime}}
// author: {{.FunctionAuthor}}
// ==========================================================================
package controller

import (
    "common/util"
    "github.com/gin-gonic/gin"
    "github.com/lostvip-com/lv_framework/utils/lv_conv"
    "{{.ModuleName}}/internal/{{.PackageName}}/service"
)

// PreList 查询页
func (w {{.ClassName}}Api) PreList{{.ClassName}}(c *gin.Context) {
	util.WriteTpl(c, "{{.PackageName}}/list_{{.BusinessName}}.html")
}
// PreAdd 新增页
func (w {{.ClassName}}Api) PreAdd{{.ClassName}}(c *gin.Context) {
	util.WriteTpl(c, "{{.PackageName}}/add_{{.BusinessName}}.html")
}
// PreEdit 修改页
func (w {{.ClassName}}Api) PreEdit{{.ClassName}}(c *gin.Context) {
	id := lv_conv.Int64(c.Query("id"))
    {{.BusinessName}}Service :=service.{{.ClassName}}Service{}
	entity, err := {{.BusinessName}}Service.FindById(id)
	if err != nil || entity == nil {
		util.WriteErrorTPL(c, gin.H{"desc": "数据不存在"})
		return
	}
	util.WriteTpl(c, "{{.PackageName}}/edit_{{.BusinessName}}.html",gin.H{
		"{{.BusinessName}}": entity,
	})
}