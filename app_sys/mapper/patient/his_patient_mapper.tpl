-- name: test
select * from his_patient where 1=1
{{if and (ne .Name "")  (ne .Name nil) }}
  and name like '%{{.Name}}%'
{{end}}
