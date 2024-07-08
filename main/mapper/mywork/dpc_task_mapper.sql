

-- name: ListByPage
select
  t.id
  , t.username
  , t.password
  , t.prj_code
  , t.task_content
  , t.start_date
  , t.end_date
  , t.work_days
  , t.auto_submit
  , t.status
  , t.sort
  , t.update_by
  , t.update_time
  , t.create_time
  , t.create_by
  , t.del_flag

from dpc_task t where 1=1 and del_flag=0

    {{if and (ne .Username "")  (ne .Username nil) }}
        and  t.username like concat('%', @username,'%')
    {{end}}
    {{if and (ne .Password "")  (ne .Password nil) }}
        and  t.password like concat('%', @password,'%')
    {{end}}
    {{if and (ne .PrjCode "")  (ne .PrjCode nil) }}
        and  t.prj_code like concat('%', @prjCode,'%')
    {{end}}
    {{if and (ne .TaskContent "")  (ne .TaskContent nil) }}
        and  t.task_content like concat('%', @taskContent,'%')
    {{end}}