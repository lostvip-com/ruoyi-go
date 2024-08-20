

-- name: ListIotDevice
select
  t.id, t.product_id , t.gateway_id,t.dept_id
  , t.node_type
  , t.name
  , t.sn
  , t.dev_no
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
  ,dept.dept_name DeptName
from iot_device t
left join sys_dept dept on dept.dept_id = t.dept_id
where 1=1 and t.del_flag=0

   {{if (ne .ProductId 0) }}
        and  t.product_id = @ProductId
    {{end}}

   {{if (ne .ProductId 0) }}
        and  t.product_id = @ProductId
    {{end}}
    {{if (ne .GatewayId 0) }}
        and  t.gateway_id = @GatewayId
    {{end}}
    {{if (ne .Name "") }}
        and  t.name like concat('%', @Name,'%')
    {{end}}
    {{if (ne .Sn "") }}
        and  t.sn like concat('%', @Sn,'%')
    {{end}}
    {{if (ne .Ancestors "") }}
        and  dept.ancestors like concat(@Ancestors,'%')
    {{end}}

    {{if (ne .NodeTypeArr nil) }}   and  t.node_type in @NodeTypeArr   {{end}}

    order by  id  desc,gateway_id desc



