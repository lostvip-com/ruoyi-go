

-- name: ListIotDriverInstance
select
     t.remark
      ,t.ping_interval
      ,t.ping_time
      ,t.platform
      ,t.driver_type
      ,t.config
      ,t.address
      ,t.name
      ,t.id

from iot_driver_instance t where 1=1 and t.del_flag=0

    {{if (ne .Address "") }}
        and  t.address like concat('%', @Address,'%')
    {{end}}
    {{if (ne .Name "") }}
        and  t.name like concat('%', @Name,'%')
    {{end}}

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



