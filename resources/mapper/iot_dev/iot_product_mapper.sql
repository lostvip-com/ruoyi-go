

-- name: ListIotProduct
select
  t.id
  , t.name
  , t.key
  , t.cloud_product_id
  , t.cloud_instance_id
  , t.platform
  , t.protocol
  , t.node_type
  , t.net_type
  , t.data_format
  , t.last_sync_time
  , t.factory
  , t.description
  , t.status
  , t.extra
  , t.del_flag
  , t.create_time
  , t.update_time
  , t.update_by
  , t.create_by

from iot_product t where 1=1 and del_flag=0

    {{if (ne .Name "") }}
        and  t.name like concat('%', @Name,'%')
    {{end}}

    {{if (ne .Test "") }}
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



