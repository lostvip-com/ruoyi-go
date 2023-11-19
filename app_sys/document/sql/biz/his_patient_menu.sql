/*
==========================================================================
LV自动生成菜单SQL,只生成一次,按需修改.
生成日期：2023-10-29 22:05:20 +0800 CST
生成路径: document/sql/biz/his_patient_menu.sql
生成人：lv
==========================================================================
*/

-- name: menu
insert into sys_menu (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('患者基本信息', '4', '1', '/biz/patient', 'C', '0', 'biz:patient:view', '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '患者基本信息菜单');

-- name: menu_button_retrieve
insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('患者基本信息查询', @parentId, '1',  '#',  'F', '0', 'biz:patient:list',         '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');
-- name: menu_button_create
insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('患者基本信息新增', @parentId, '2',  '#',  'F', '0', 'biz:patient:add',          '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');
-- name: menu_button_update
insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('患者基本信息修改', @parentId, '3',  '#',  'F', '0', 'biz:patient:edit',         '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');
-- name: menu_button_delete
insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('患者基本信息删除', @parentId, '4',  '#',  'F', '0', 'biz:patient:remove',       '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');
-- name: menu_button_export
insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('患者基本信息导出', @parentId, '5',  '#',  'F', '0', 'biz:patient:export',       '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');