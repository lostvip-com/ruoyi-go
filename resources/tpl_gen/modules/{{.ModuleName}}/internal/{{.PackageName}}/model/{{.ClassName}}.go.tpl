// ==========================================================================
// LV自动生成数据库操作代码，无需手动修改，重新生成会自动覆盖.
// 生成日期: {{.CreateTime}}
// 生成人: {{.FunctionAuthor}}
// ==========================================================================
package model

import (
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/namedsql"
	"time"
)

// {{.ClassName}} {{.TableComment}}
type {{.ClassName}} struct {
{{- range $index, $column := .Columns -}}

{{- if eq $column.IsPk "1" }}
    {{$column.GoField}}  {{$column.GoType}}  `gorm:"type:{{$column.ColumnType}};primary_key;auto_increment;{{$column.ColumnComment}};" json:"{{$column.HtmlField}}"`
{{- end -}}
{{- if ne $column.IsPk "1" }}
    {{- if eq $column.ColumnType "datetime" }}
    {{$column.GoField}} time.Time `gorm:"type:datetime;comment:{{$column.ColumnComment}};" json:"{{$column.HtmlField}}" time_format:"2006-01-02 15:04:05"`
    {{- else if eq $column.ColumnType "date" }}
    {{$column.GoField}} time.Time `gorm:"type:date;comment:{{$column.ColumnComment}};" json:"{{$column.HtmlField}}" time_format:"2006-01-02"`
    {{- else }}
    {{$column.GoField}}  {{$column.GoType}}  `gorm:"type:{{$column.ColumnType}};comment:{{$column.ColumnComment}};" json:"{{$column.HtmlField}}"`
    {{- end -}}
{{- end -}}
{{- end }}
   
}

func (e *{{.ClassName}}) TableName() string {
	return "{{.TbName}}"
}

func (e *{{.ClassName}}) Save() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

func (e *{{.ClassName}}) FindById(id int64) (*{{.ClassName}},error) {
	err := lv_db.GetMasterGorm().Take(e,id).Error
	return e,err
}

func (e *{{.ClassName}}) FindOne() (*{{.ClassName}},error) {
    tb := lv_db.GetMasterGorm().Table(e.TableName())
{{range $index, $column := .Columns -}}
{{if eq $column.IsQuery "0"}}
    {{- continue -}}
{{- end -}}
{{ if contains $column.GoType "int"}}
    if e.{{$column.GoField}} != 0 {
        tb = tb.Where("{{$column.ColumnName}}=?", e.{{$column.GoField}})
    }
{{- end -}}
{{ if eq $column.GoType "string" }}
    if e.{{$column.GoField}} != "" {
         tb = tb.Where("{{$column.ColumnName}}=?", e.{{$column.GoField}})
    }
{{- end -}}
{{- end }}
    err := tb.First(e).Error
    return e,err
}

func (e *{{.ClassName}}) Updates() error {
	return lv_db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}

func (e *{{.ClassName}}) Delete() error {
	return lv_db.GetMasterGorm().Delete(e).Error
}

func (e *{{.ClassName}}) Count() (int64, error) {
	sql := " select count(*) from {{.TbName}} where del_flag = 0 "

	{{range $index, $column := .Columns -}}
    {{if eq $column.IsQuery "0"}}
        {{- continue -}}
    {{- end -}}
    {{ if contains $column.GoType "int"}}
        if e.{{$column.GoField}} != 0 {
            sql += " and {{$column.ColumnName}} = @{{$column.GoField}} "
        }
    {{- end -}}
    {{ if eq $column.GoType "string" }}
         if e.{{$column.GoField}} != "" {
            sql += " and {{$column.ColumnName}} = @{{$column.GoField}} "
         }
    {{- end -}}
    {{- end }}

	return namedsql.Count(lv_db.GetMasterGorm(), sql, e)
}