-- name: ListByPage
select
  t.id id
  , t.username Username
  , t.password password
  , t.prj_code
  , t.task_content
  , t.work_days workDays
  , t.auto_submit autoSubmit
  , t.status status
  , t.sort sort
  , t.del_flag delFlag
from dpc_task t where del_flag=0
{{if and (ne .Username "")  (ne .Username nil) }}
    and  t.username like concat('%',@username,'%')
{{end}}
{{if and (ne .PrjCode "")  (ne .PrjCode nil) }}
    and  t.prj_code = @prjCode
{{end}}
{{if and (ne .TaskContent "")  (ne .TaskContent nil) }}
    and  t.task_content like concat('%',@taskContent,'%')
{{end}}

-- name: ListAll
select
  t.id id
  , t.username Username
  , t.password password
  , t.prj_code
  , t.task_content
  , t.work_days workDays
  , t.auto_submit autoSubmit
  , t.status status
  , t.sort sort
  , t.del_flag delFlag
from dpc_task t where del_flag=0
{{if and (ne .Username "")  (ne .Username nil) }}
    and  t.username like concat('%',@username,'%')
{{end}}
{{if and (ne .PrjCode "")  (ne .PrjCode nil) }}
    and  t.prj_code = @prjCode
{{end}}
{{if and (ne .TaskContent "")  (ne .TaskContent nil) }}
    and  t.task_content like concat('%',@taskContent,'%')
{{end}}