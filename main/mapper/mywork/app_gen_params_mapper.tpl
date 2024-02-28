

-- name: ListByPage
select
  t.id
  , t.use_flag
  , t.param_no
  , t.param_name
  , t.param_type
  , t.unit
  , t.remark
  , t.monitor_type_id
  , t.create_time

from app_gen_params t where 1=1 and del_flag=0

   {{if and (ne .ParamNo 0)  (ne .ParamNo nil) }}
        and  t.param_no = @ParamNo
    {{end}}
    {{if and (ne .ParamName "")  (ne .ParamName nil) }}
        and  t.param_name like concat('%', @ParamName,'%')
    {{end}}
    {{if and (ne .ParamType "")  (ne .ParamType nil) }}
        and  t.param_type = @ParamType
    {{end}}