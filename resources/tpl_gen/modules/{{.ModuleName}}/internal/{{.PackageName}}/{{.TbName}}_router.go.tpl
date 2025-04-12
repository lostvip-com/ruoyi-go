// ==========================================================================
// LV自动生成路由代码,只生成一次,按需修改,再次生成不会覆盖.
// 生成日期:{{.CreateTime}}
// 生成人:{{.FunctionAuthor}}
// ==========================================================================
package {{.PackageName}}

import (
        "github.com/lostvip-com/lv_framework/web/router"
        "common/middleware/auth"
        "{{.ModuleName}}/internal/{{.PackageName}}/controller"
)

func init() {
	group_{{.BusinessName}} := router.New( "/{{.PackageName}}/{{.BusinessName}}", auth.TokenCheck())

	{{.BusinessName}} := controller.{{.ClassName}}Api{}
	group_{{.BusinessName}}.GET("/", "{{.PackageName}}:{{.BusinessName}}:view", {{.BusinessName}}.PreList{{.ClassName}})
	group_{{.BusinessName}}.GET("/preAdd{{.ClassName}}", "{{.PackageName}}:{{.BusinessName}}:add", {{.BusinessName}}.PreAdd{{.ClassName}})
	group_{{.BusinessName}}.GET("/preEdit{{.ClassName}}", "{{.PackageName}}:{{.BusinessName}}:edit", {{.BusinessName}}.PreEdit{{.ClassName}})
	// api
	group_{{.BusinessName}}.POST("/list{{.ClassName}}", "{{.PackageName}}:{{.BusinessName}}:list", {{.BusinessName}}.List{{.ClassName}})
	group_{{.BusinessName}}.POST("/add{{.ClassName}}", "{{.PackageName}}:{{.BusinessName}}:add", {{.BusinessName}}.Add{{.ClassName}})
	group_{{.BusinessName}}.POST("/remove{{.ClassName}}", "{{.PackageName}}:{{.BusinessName}}:remove", {{.BusinessName}}.Remove{{.ClassName}})
	group_{{.BusinessName}}.POST("/edit{{.ClassName}}", "{{.PackageName}}:{{.BusinessName}}:edit",{{.BusinessName}}.Save{{.ClassName}})
	group_{{.BusinessName}}.POST("/export{{.ClassName}}", "{{.PackageName}}:{{.BusinessName}}:export", {{.BusinessName}}.Export{{.ClassName}})
}
