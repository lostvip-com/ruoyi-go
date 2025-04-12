

-- name: ListByPage
select
  t.id id
  , t.task_id taskId
  , t.action action
  , t.content content
  , t.idx_xpath idxXpath
  , t.idx_img1 idxImg1
  , t.idx_img2 idxImg2
  , t.idx_img3 idxImg3
  , t.status status
  , t.sort sort
  , t.del_flag delFlag
  , t.create_by createBy
  , t.create_time createTime
  , t.update_by updateBy
  , t.update_time updateTime

from dpc_task_item t where 1=1 and del_flag=0

   {{if and (ne .TaskId 0)  (ne .TaskId nil) }}
        and  t.task_id = @taskId
    {{end}}
    {{if and (ne .Action "")  (ne .Action nil) }}
        and  t.action like concat('%', @action,'%')
    {{end}}
    {{if and (ne .Content "")  (ne .Content nil) }}
        and  t.content like concat('%', @content,'%')
    {{end}}
    {{if and (ne .IdxXpath "")  (ne .IdxXpath nil) }}
        and  t.idx_xpath like concat('%', @idxXpath,'%')
    {{end}}