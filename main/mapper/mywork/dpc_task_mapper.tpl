

-- name: ListByPage
select
  t.id id
  , t.username username
  , t.password password
  , t.prj_code prjCode
  , t.task_content taskContent
  , t.start_date startDate
  , t.end_date endDate
  , t.work_days workDays
  , t.auto_submit autoSubmit
  , t.status status
  , t.sort sort
  , t.del_flag delFlag
  , t.create_by createBy
  , t.update_by updateBy
  , t.create_time createTime
  , t.update_time updateTime

from dpc_task t where 1=1 and del_flag=0

    {{if and (ne .Username "")  (ne .Username nil) }}
        and  t.username = @username
    {{end}}
    {{if and (ne .Password "")  (ne .Password nil) }}
        and  t.password = @password
    {{end}}
    {{if and (ne .PrjCode "")  (ne .PrjCode nil) }}
        and  t.prj_code = @prjCode
    {{end}}
    {{if and (ne .TaskContent "")  (ne .TaskContent nil) }}
        and  t.task_content = @taskContent
    {{end}}