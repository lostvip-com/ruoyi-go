

-- name: ListAppGenParams
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
        and  t.param_no like concat('%',@ParamNo,'%')
    {{end}}
    {{if and (ne .ParamName "")  (ne .ParamName nil) }}
        and  t.param_name like concat('%', @ParamName,'%')
    {{end}}
    {{if and (ne .ParamType "")  (ne .ParamType nil) }}
        and  t.param_type = @ParamType
    {{end}}
   {{if and (ne .MonitorTypeId 0)  (ne .MonitorTypeId nil) }}
        and  t.monitor_type_id = @MonitorTypeId
    {{end}}