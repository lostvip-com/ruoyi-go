-- name: ListIotPrdEvent
select
    t.id
     , t.product_id
     , t.event_type
     , t.code
     , t.tag
     , t.name
     , t.remark
     , t.output_params
     , t.del_flag
     , t.create_time
     , t.update_time
     , t.update_by
     , t.create_by
     , t.tenant_id

from iot_prd_event t where 1=1 and t.del_flag=0

    {{if (ne .ProductId 0) }}
        and  t.product_id = @ProductId
    {{end}}
    {{if (ne .EventType "") }}
        and  t.event_type like concat('%', @EventType,'%')
    {{end}}
    {{if (ne .Code "") }}
        and  t.code like concat('%', @Code,'%')
    {{end}}
    {{if (ne .Tag "") }}
        and  t.tag like concat('%', @Tag,'%')
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