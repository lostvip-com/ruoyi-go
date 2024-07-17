

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

    {{if (ne .ParamNoStart 0) }}
        and  t.param_no >= @ParamNoStart
    {{end}}
    {{if (ne .ParamNoEnd 0) }}
            and  t.param_no <= @ParamNoEnd
    {{end}}
    {{if and (ne .ParamName "")  (ne .ParamName nil) }}
        and  t.param_name like concat('%', @ParamName,'%')
    {{end}}
    {{if (ne .Remark "") }}
        and  t.remark like concat('%', @Remark,'%')
    {{end}}
    {{if and (ne .ParamType "")  (ne .ParamType nil) }}
        and  t.param_type = @ParamType
    {{end}}
    {{if and (ne .MonitorTypeId 0)  (ne .MonitorTypeId nil) }}
        and  t.monitor_type_id = @MonitorTypeId
    {{end}}
    {{if and (ne .UseFlag "")  (ne .UseFlag nil) }}
        and  t.use_flag = @UseFlag
    {{end}}

     order by
     {{if (eq .SortName "paramNo") }}
        param_no
     {{else}}
        id
     {{end}}

     {{if (eq .SortOrder "asc") }}
          asc
     {{else}}
          desc
     {{end}}

