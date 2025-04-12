

-- name: ListIotDataHis
select
     t.value
      ,t.name
      ,t.tags
      ,t.tenant_id
      ,t.create_by
      ,t.update_by
      ,t.update_time
      ,t.create_time
      ,t.del_flag
      ,t.dev_time
      ,t.code
      ,t.device_id
      ,t.id

from iot_data_his t where 1=1 and t.del_flag=0

    {{if (ne .Name "") }}
        and  t.name like concat('%', @Name,'%')
    {{end}}
    {{if (ne .TenantId "") }}
        and  t.tenant_id like concat('%', @TenantId,'%')
    {{end}}
    {{if (ne .Code "") }}
        and  t.code like concat('%', @Code,'%')
    {{end}}
   {{if (ne .DeviceId 0) }}
        and  t.device_id = @DeviceId
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



