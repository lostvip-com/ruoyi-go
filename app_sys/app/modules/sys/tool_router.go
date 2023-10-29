package sys

import (
	"lostvip.com/router"
	"robvi/app/middleware/auth"
	"robvi/app/middleware/jwt"
	"robvi/app/modules/sys/controller/tool"
)

// 加载路由
func init() {
	// 服务监控
	g1 := router.New("/tool", jwt.JWTAuthMiddleware(), auth.Auth)
	g1.GET("/build", "tool:build:view", tool.Build)
	g1.GET("/swagger", "tool:swagger:view", tool.Swagger)
	g1.GET("/gen", "tool:gen:view", tool.Gen)
	g1.POST("/gen/list", "tool:gen:list", tool.GenList)
	g1.POST("/gen/remove", "tool:gen:remove", tool.Remove)
	g1.GET("/gen/importTable", "tool:gen:list", tool.ImportTable)
	g1.POST("/gen/db/list", "tool:gen:list", tool.DataList)
	g1.POST("/gen/importTable", "tool:gen:list", tool.ImportTableSave)
	g1.GET("/gen/edit", "tool:gen:edit", tool.Edit)
	g1.POST("/gen/edit", "tool:gen:edit", tool.EditSave)
	g1.POST("/gen/column/list", "tool:gen:list", tool.ColumnList)
	g1.GET("/gen/preview", "tool:gen:preview", tool.Preview)
	g1.GET("/gen/genCode", "tool:gen:code", tool.GenCode)
}
