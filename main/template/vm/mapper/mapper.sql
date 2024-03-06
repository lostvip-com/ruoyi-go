{{$tagS := "{{"}}
{{$tagE := "}}"}}
-- name: List{{.table.ClassName}}
select
 {{- range $index, $column := .table.Columns}}
 {{if ne $index 0}} ,{{end}} t.{{$column.ColumnName}}
 {{- end }}

from {{.table.TbName}} t where 1=1 and del_flag=0
{{range $index, $column := .table.Columns -}}
{{if eq $column.IsQuery "0"}}
    {{- continue -}}
{{- end -}}
{{- if contains $column.GoType "int"}}
   {{$tagS}}if (ne .{{$column.GoField}} 0) {{$tagE}}
        and  t.{{$column.ColumnName}} = @{{$column.GoField}}
    {{$tagS}}end{{$tagE}}
{{- end -}}
{{if eq $column.GoType "string"}}
    {{$tagS}}if (ne .{{$column.GoField}} "") {{$tagE}}
        {{- if eq $column.QueryType "LIKE" }}
        and  t.{{$column.ColumnName}} like concat('%', @{{$column.GoField}},'%')
        {{- else}}
        and  t.{{$column.ColumnName}} = @{{$column.GoField}}
        {{- end }}
    {{$tagS}}end{{$tagE}}
    {{- end -}}
{{- end }}

    order by
        {{$tagS}}if (eq .SortName "{{.table.PkColumn.HtmlField}}") {{$tagE}}
           {{.table.PkColumn.ColumnName}}
        {{$tagS}}else{{$tagE}}
           {{.table.PkColumn.ColumnName}}
        {{$tagS}}end{{$tagE}}

        {{$tagS}}if (eq .SortOrder "asc") {{$tagE}}
             asc
        {{$tagS}}else{{$tagE}}
             desc
        {{$tagS}}end{{$tagE}}



