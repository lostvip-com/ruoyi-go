

-- name: ListIotDevice
select
  t.id
  , t.product_id
  , t.name
  , t.status
  , t.description
  , t.secret
  , t.platform
  , t.install_location
  , t.last_sync_time
  , t.last_online_time
  , t.del_flag
  , t.create_time
  , t.update_time
  , t.update_by
  , t.create_by

from iot_device t where 1=1 and del_flag=0

   {{if (ne .ProductId 0) }}
        and  t.product_id = @ProductId
    {{end}}

   {{if (ne .ProductId 0) }}
        and  t.product_id = @ProductId
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



