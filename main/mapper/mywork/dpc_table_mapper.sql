

-- name: ListByPage
select
  t.id id
  , t.username username
  , t.password password
  , t.taskContent taskContent
  , t.startDate startDate
  , t.endDate endDate
  , t.workDays workDays
  , t.autoSubmit autoSubmit
  , t.status status
  , t.sort sort
  , t.updateTime updateTime

from dpc_table t where 1=1 and del_flag=0

    {{if and (ne .Username "")  (ne .Username nil) }}
        and  t.username like concat('%', @username,'%')
    {{end}}
    {{if and (ne .Password "")  (ne .Password nil) }}
        and  t.password like concat('%', @password,'%')
    {{end}}
    {{if and (ne .TaskContent "")  (ne .TaskContent nil) }}
        and  t.taskContent like concat('%', @taskContent,'%')
    {{end}}