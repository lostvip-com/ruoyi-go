

-- name: ListIotScene
select
  t.tenant_id
  , t.create_by
  , t.update_by
  , t.update_time
  , t.create_time
  , t.del_flag
  , t.remark
  , t.actions
  , t.conditions
  , t.status
  , t.name
  , t.id

from iot_scene t where 1=1 and t.del_flag=0

    {{if (ne .Status "") }}
        and  t.status = @Status
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



