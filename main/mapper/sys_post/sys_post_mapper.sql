-- name: listSql
select * from sys_post p where 1=1
{{if and (ne .PostName "")  (ne .PostName nil) }}
  and post_name like concat('%',@PostName,'%')
{{end}}
{{if and (ne .Remark "")  (ne .Remark nil) }}
  and remark like @Remark
{{end}}

-- name: listSql2
select post_id,post_name,remark from sys_post p where 1=1
{{if and (ne .PostName "")  (ne .PostName nil) }}
  and post_name like concat('%',@PostName,'%')
{{end}}

{{if and (ne .Remark "")  (ne .Remark nil) }}
  and remark like @Remark
{{end}}


