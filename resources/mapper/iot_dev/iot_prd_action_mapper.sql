-- name: ListIotPrdAction
select
    t.id
     , t.product_id
     , t.code
     , t.tag
     , t.name
     , t.call_type
     , t.input_params
     , t.output_params
     , t.remark
     , t.del_flag
     , t.create_time
     , t.update_time
     , t.update_by
     , t.create_by
     , t.tenant_id

from iot_prd_action t where 1=1 and t.del_flag=0

    {{if (ne .ProductId 0) }}
        and  t.product_id = @ProductId
    {{end}}
    {{if (ne .Code "") }}
        and  t.code like concat('%', @Code,'%')
    {{end}}
    {{if (ne .Tag "") }}
        and  t.tag like concat('%', @Tag,'%')
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