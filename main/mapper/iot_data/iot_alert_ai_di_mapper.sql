

-- name: ListIotAlertAiDi
select
  t.tenant_id
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
  , t.dev_id
  , t.id

from iot_alert_ai_di t where 1=1 and t.del_flag=0


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



