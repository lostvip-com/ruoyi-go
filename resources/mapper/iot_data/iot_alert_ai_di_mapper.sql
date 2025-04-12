

-- name: ListIotAlertAiDi
select
  dev.name as device_name
  , t.create_by
  , t.update_by
  , t.update_time
  , t.create_time
  , t.del_flag
  , t.status
  , t.treated_time
  , t.recover_time
  , t.trigger_time
  , t.message
  , t.content
  , t.level
  , t.type
  , t.code
  , t.device_id
  , t.dept_id
  , t.id

from      iot_data_event t
left join iot_device dev on t.device_id = dev.id
    {{if (ne .Ancestors "") }}
        left join sys_dept dept on t.dept_id = dept.dept_id
    {{end}}

    where 1=1 and t.del_flag=0

    {{if (ne .Ancestors "") }}
        and dept.ancestors like @Ancestors
    {{end}}

     {{if (ne .DeviceId 0) }}  and  t.device_id = @DeviceId {{end}}
    order by
        {{if (eq .SortName "id") }}
           id
        {{else}}
           id
        {{end}}

        {{if (eq .SortOrder "asc") }}
             asc
        {{else}}
             desc
        {{end}}



