{{$tagS := "{{"}}
{{$tagE := "}}"}}
-- name: ListByPage
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
   {{$tagS}}if and (ne .{{$column.GoField}} 0)  (ne .{{$column.GoField}} nil) {{$tagE}}
        and  t.{{$column.ColumnName}} = @{{$column.GoField}}
    {{$tagS}}end{{$tagE}}
{{- end -}}
{{if eq $column.GoType "string"}}
    {{$tagS}}if and (ne .{{$column.GoField}} "")  (ne .{{$column.GoField}} nil) {{$tagE}}
        {{- if eq $column.QueryType "LIKE" }}
        and  t.{{$column.ColumnName}} like concat('%', @{{$column.GoField}},'%')
        {{- else}}
        and  t.{{$column.ColumnName}} = @{{$column.GoField}}
        {{- end }}
    {{$tagS}}end{{$tagE}}
    {{- end -}}
{{- end -}}


