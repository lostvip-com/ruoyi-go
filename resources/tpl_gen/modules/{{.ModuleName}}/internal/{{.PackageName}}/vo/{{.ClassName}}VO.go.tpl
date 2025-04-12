// ==========================================================================
// LV自动生成model扩展代码列表 按需修改
// 生成日期: {{.CreateTime}}
// 生成人: {{.FunctionAuthor}}
// ==========================================================================
package vo

import (
  "github.com/lostvip-com/lv_framework/web/lv_dto"
  "time"
)

//新增页面请求参数
type Add{{.ClassName}}Req struct {
{{- range $index, $column := .Columns}}
    {{- if and (eq $column.IsInsert "1") (ne $column.IsPk "1")}}
    {{  $column.GoField}} {{$column.GoType}} `form:"{{$column.HtmlField}}"
    {{- if eq $column.IsRequired "1"}} binding:"required" {{end -}}
    {{- if eq $column.ColumnType "date"}} time_format:"2006-01-02" {{end}}
    {{- if eq $column.ColumnType "datetime"}} time_format:"2006-01-02 15:04:05" {{end}}`
    {{- end}}
{{- end}}
    CreateBy string
}

//修改页面请求参数
type Edit{{.ClassName}}Req struct {
    {{.PkColumn.GoField}}  {{.PkColumn.GoType}}  `form:"{{.PkColumn.HtmlField}}" binding:"required"`
  {{ range $index, $column := .Columns}}
    {{- if and (eq $column.IsInsert "1") (ne $column.IsPk "1")}}
    {{  $column.GoField}} {{$column.GoType}} `form:"{{$column.HtmlField}}"
    {{- if eq $column.IsRequired "1"}}binding:"required" {{end}}
    {{- if eq $column.ColumnType "date"}} time_format:"2006-01-02" {{end}}
    {{- if eq $column.ColumnType "datetime"}} time_format:"2006-01-02 15:04:05" {{end}}`
   	{{- end }}
{{- end}}
    UpdateBy string
}

//分页请求参数 {{$pkColumn := .PkColumn}}
type {{.ClassName}}Req struct {
{{- range $index, $column := .Columns }}
{{- if ne $column.IsQuery "1" -}}
	{{- continue -}}
{{- end -}}
{{if eq $column.IsQuery "1" }}
  {{$column.GoField}} {{$column.GoType}} `form:"{{$column.HtmlField}}"` //{{$column.ColumnComment}}
{{- end -}}
{{- end}}
    BeginTime  string `form:"beginTime"`  //开始时间
    EndTime    string `form:"endTime"`    //结束时间
    lv_dto.Paging
}


//分页请求结果
type {{.ClassName}}Resp struct {
  {{.PkColumn.GoField}}    {{.PkColumn.GoType}}  `json:"{{.PkColumn.HtmlField}}"`
  {{- range $index, $column := .Columns}}
    {{- if and (eq $column.IsInsert "1") (ne $column.IsPk "1")}}
        {{ $column.GoField}} {{$column.GoType}} `json:"{{$column.HtmlField}}"
        {{- if eq $column.ColumnType "date"}} time_format:"2006-01-02" {{end}}
        {{- if eq $column.ColumnType "datetime"}} time_format:"2006-01-02 15:04:05" {{end}}`
    {{- end }}
  {{- end}}
  UpdateBy string      `json:"updateBy"`
  UpdateTime time.Time `json:"updateTime" time_format:"2006-01-02 15:04:05"`
  CreateTime time.Time `json:"createTime" time_format:"2006-01-02 15:04:05"`
  CreateBy   string    `json:"createBy"`
}
