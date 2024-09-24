

-- name: ListIotDataHis
select
  t.tenant_id
  , t.create_by
  , t.update_by
  , t.update_time
  , t.create_time
  , t.del_flag
  , t.dev_time
  , t.prop_value
  , t.prop_name
  , t.prop_tags
  , t.code
  , t.dev_id
  , t.id

from iot_data_his t where 1=1 and t.del_flag=0


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



