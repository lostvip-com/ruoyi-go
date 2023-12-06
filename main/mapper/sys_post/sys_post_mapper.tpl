-- name: test
select * from sys_post where 1=1
{{if and (ne .PostName "")  (ne .PostName nil) }}
  and post_name like '%{{.PostName}}%'
{{end}}
