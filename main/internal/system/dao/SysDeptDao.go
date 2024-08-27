package dao

import (
	"common/model"
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/lv_db/lv_dao"
	"github.com/spf13/cast"
)

// Fill with you ideas below.

// GenTable is the golang structure for table sys_dept.
type SysDeptDao struct {
}

// 根据部门ID查询信息
func (dao SysDeptDao) SelectDeptById(deptId int64) (*model.SysDept, error) {
	var result = new(model.SysDept)
	result, err := result.FindById(deptId)
	return result, err
}

// 根据ID查询所有子部门
func (dao SysDeptDao) SelectChildrenDeptById(deptId int64) []*model.SysDept {
	db := lv_db.GetMasterGorm()
	var rs []*model.SysDept
	//lv_db.Table("sys_dept").Where("find_in_set(?, ancestors)", deptId).Find(&rs)
	db.Table("sys_dept").Where("parent_id=?", deptId).Find(&rs)
	return rs
}

// 删除部门管理信息
func (dao SysDeptDao) DeleteDeptById(deptId int64) error {
	var entity model.SysDept
	err := entity.UpdateDelFlag(deptId)
	return err
}

// 修改子元素关系（替换前半部分）
func (dao SysDeptDao) UpdateDeptChildrenAncestors(dept *model.SysDept, parentCodes string) {
	dept.Ancestors = parentCodes + "," + cast.ToString(dept.DeptId)
	lv_db.GetMasterGorm().Table("sys_dept").Where("dept_id=", dept.DeptId).Update("ancestors", dept.Ancestors)
	// ancestors 上级ancestors发生变化，修改下级
	deptList := dao.SelectChildrenDeptById(dept.DeptId)
	if deptList == nil || len(deptList) <= 0 {
		return
	}
	for _, child := range deptList {
		dao.UpdateDeptChildrenAncestors(child, dept.Ancestors)
	}
}

// 查询部门管理数据
func (d SysDeptDao) SelectDeptList(parentId int64, deptName, status string, tenantId int64) (*[]model.SysDept, error) {
	sql := ` select *  from sys_dept d  where d.del_flag =0  `
	param := map[string]any{}
	if parentId > 0 {
		param["parentId"] = parentId
		sql += " and d.parent_id = @parentId "
	}
	if deptName != "" {
		param["deptName"] = "%" + deptName + "%"
		sql += " and d.dept_name like @deptName "
	}
	if status != "" {
		param["status"] = status
		sql += " and d.status = @status "
	}
	if tenantId != 0 {
		param["tenantId"] = tenantId
		sql += " and d.tenant_id = @tenantId "
	}
	sql += " order by d.parent_id, d.order_num "

	return lv_dao.ListByNamedSql[model.SysDept](sql, param)
}

// 根据角色ID查询部门
func (dao SysDeptDao) SelectRoleDeptTree(roleId int64) ([]string, error) {
	sql := ` select concat(d.dept_id, d.dept_name) as DeptName 
             from sys_dept d 
             left join sys_role_dept rd  on d.dept_id = rd.dept_id 
             where d.del_flag =0 and rd.role_id = @roleId
             order by d.parent_id, d.order_num
             `
	param := map[string]any{}
	param["roleId"] = roleId
	listMap, err := lv_dao.ListMapStrByNamedSql(sql, param, false)
	var result []string
	var rs = *listMap
	if err == nil && rs != nil && len(rs) > 0 {
		for _, record := range rs {
			if record["DeptName"] != "" {
				result = append(result, record["DeptName"])
			}
		}
	}
	return result, nil
}

// 查询部门是否存在用户
func (dao SysDeptDao) CheckDeptExistUser(deptId int64) bool {
	sql := " select count(*) from sys_user where del_flag = 0 "
	param := map[string]any{}
	param["deptId"] = deptId
	sql += " and dept_id= @deptId "
	count, err := lv_dao.CountByNamedSql(sql, param)
	if err != nil {
		panic(err)
	}
	if count > 0 {
		return true
	} else {
		return false
	}
}

// 查询部门人数
func (dao SysDeptDao) SelectDeptCount(deptId, parentId int64) int64 {
	sql := " select count(*) from sys_dept where del_flag = 0 "
	param := map[string]any{}
	if deptId > 0 {
		param["deptId"] = deptId
		sql += " and dept_id= @deptId "
	}
	if parentId > 0 {
		param["parentId"] = parentId
		sql += " and parent_id= @parentId "
	}
	count, err := lv_dao.CountByNamedSql(sql, param)
	if err != nil {
		panic(err)
	}
	return count
}

// 校验部门名称是否唯一
func (dao SysDeptDao) FindOne(deptName string, parentId int64) (*model.SysDept, error) {
	var entity model.SysDept
	entity.DeptName = deptName
	entity.ParentId = parentId
	err := entity.FindOne()
	return &entity, err
}
