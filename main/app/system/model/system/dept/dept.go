package dept

import (
	"errors"
	"fmt"
	"lostvip.com/db"
	"lostvip.com/utils/lv_conv"
	"strings"
)

// Fill with you ideas below.

// GenTable is the golang structure for table sys_dept.
type SysDeptExtend struct {
	SysDept    `xorm:"extends"`
	ParentName string `json:"parentName"`
}

// 分页请求参数
type SelectPageReq struct {
	ParentId  int64  `form:"parentId"`      //父部门ID
	DeptName  string `form:"deptName"`      //部门名称
	Status    string `form:"status"`        //状态
	PageNum   int    `form:"pageNum"`       //当前页码
	PageSize  int    `form:"pageSize"`      //每页数
	SortName  string `form:"orderByColumn"` //排序字段
	SortOrder string `form:"isAsc"`         //排序方式
	BeginTime string `form:"beginTime"`     //开始时间
	EndTime   string `form:"endTime"`       //结束时间
	TenantId  int64  `form:"tenantId"`      //结束时间
}

// 新增页面请求参数
type AddReq struct {
	ParentId int64  `form:"parentId"  binding:"required"`
	DeptName string `form:"deptName"  binding:"required"`
	OrderNum int    `form:"orderNum" binding:"required"`
	Leader   string `form:"leader"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
	Status   string `form:"status"`
	TenantId int64  `form:"tenantId"` //结束时间
}

// 修改页面请求参数
type EditReq struct {
	DeptId int64 `form:"deptId" binding:"required"`
	AddReq
}

// 检查菜单名称请求参数
type CheckDeptNameReq struct {
	DeptId   int64  `form:"deptId"  binding:"required"`
	ParentId int64  `form:"parentId"  binding:"required"`
	DeptName string `form:"deptName"  binding:"required"`
}

// 检查菜单名称请求参数
type CheckDeptNameALLReq struct {
	ParentId int64  `form:"parentId"  binding:"required"`
	DeptName string `form:"deptName"  binding:"required"`
}

// 根据部门ID查询信息
func SelectDeptById(id int64) (*SysDeptExtend, error) {
	db := db.GetInstance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	var result SysDeptExtend
	model := db.Table(TableName()).Alias("d")
	model.Select("d.dept_id, d.parent_id, d.ancestors, d.dept_name, d.order_num, d.leader, d.phone, d.email, d.status,(select dept_name from sys_dept where dept_id = d.parent_id) parent_name")
	model.Where("d.dept_id = ?", id)
	_, err := model.Get(&result)
	return &result, err
}

// 根据ID查询所有子部门
func SelectChildrenDeptById(deptId int64) []*SysDept {
	db := db.GetInstance().Engine()

	if db == nil {
		return nil
	}
	var rs []*SysDept
	db.Table(TableName()).Where("find_in_set(?, ancestors)", deptId).Find(&rs)
	return rs
}

// 删除部门管理信息
func DeleteDeptById(deptId int64) int64 {
	var entity SysDept
	entity.DeptId = deptId
	entity.DelFlag = "2"
	rs, err := entity.Update()
	if err != nil {
		return 0
	}
	return rs
}

// 修改子元素关系（替换前半部分）
func UpdateDeptChildren(deptId int64, newAncestors, oldAncestors string) {
	deptList := SelectChildrenDeptById(deptId)

	if deptList == nil || len(deptList) <= 0 {
		return
	}

	for _, tmp := range deptList {
		tmp.Ancestors = strings.ReplaceAll(tmp.Ancestors, oldAncestors, newAncestors)
	}

	ancestors := " case dept_id"
	idStr := ""

	for _, dept := range deptList {
		ancestors += " when " + lv_conv.String(dept.DeptId) + " then " + dept.Ancestors
		if idStr == "" {
			idStr = lv_conv.String(dept.DeptId)
		} else {
			idStr += "," + lv_conv.String(dept.DeptId)
		}
	}

	ancestors += " end "
	db := db.GetInstance().Engine()

	if db == nil {
		return
	}

	rs, err := db.Table(TableName()).Where("dept_id in(?)", deptId).Update(map[string]interface{}{"ancestors": ancestors})
	fmt.Printf("修改了%v行 错误信息：%v", rs, err.Error())
}

// 查询部门管理数据
func SelectDeptList(parentId int64, deptName, status string, tenantId int64) ([]SysDept, error) {
	var result []SysDept
	db := db.GetInstance().Engine()
	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	model := db.Table(TableName()).Alias("d").Where("d.del_flag = '0'")
	if parentId > 0 {
		model.Where("d.parent_id = ?", parentId)
	}
	if deptName != "" {
		model.Where("d.dept_name like ?", "%"+deptName+"%")
	}
	if status != "" {
		model.Where("d.status = ?", status)
	}
	if tenantId != 0 {
		model.Where("d.tenant_id = ?", tenantId)
	}

	model.OrderBy("d.parent_id, d.order_num")

	err := model.Find(&result)

	return result, err
}

// 根据角色ID查询部门
func SelectRoleDeptTree(roleId int64) ([]string, error) {
	db := db.GetInstance().Engine()
	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	model := db.Table(TableName()).Alias("d").Join("LEFT", []string{"sys_role_dept", "rd"}, "d.dept_id = rd.dept_id")
	model.Where("d.del_flag = '0'").Where("rd.role_id = ?", roleId)
	model.OrderBy("d.parent_id, d.order_num ")
	model.Select("concat(d.dept_id, d.dept_name) as dept_name")

	var result []string
	var rs []SysDept
	err := model.Find(&result)
	if err == nil && rs != nil && len(rs) > 0 {
		for _, record := range rs {
			if record.DeptName != "" {
				result = append(result, record.DeptName)
			}
		}
	}
	return result, nil
}

// 查询部门是否存在用户
func CheckDeptExistUser(deptId int64) bool {
	db := db.GetInstance().Engine()
	if db == nil {
		return false
	}

	num, _ := db.Table(TableName()).Where("dept_id = ? and del_flag = '0'", deptId).Count()

	if num > 0 {
		return true
	} else {
		return false
	}
}

// 查询部门人数
func SelectDeptCount(deptId, parentId int64) int64 {
	db := db.GetInstance().Engine()
	if db == nil {
		return 0
	}

	result := int64(0)
	whereStr := "del_flag = '0'"
	if deptId > 0 {
		whereStr = whereStr + " and dept_id=" + lv_conv.String(deptId)
	}
	if parentId > 0 {
		whereStr = whereStr + " and parent_id=" + lv_conv.String(parentId)
	}

	rs, err := db.Table(TableName()).Where(whereStr).Count()
	if err != nil {
		result = rs
	}
	return result
}

// 校验部门名称是否唯一
func CheckDeptNameUniqueAll(deptName string, parentId int64) (*SysDept, error) {
	var entity SysDept
	entity.DeptName = deptName
	entity.ParentId = parentId
	entity.DelFlag = "0"
	ok, err := entity.FindOne()
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}

// 根据条件查询
func Find(where, order string) ([]SysDept, error) {
	var list []SysDept
	err := db.GetInstance().Engine().Table(TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

// 指定字段集合查询
func FindIn(column string, args ...interface{}) ([]SysDept, error) {
	var list []SysDept
	err := db.GetInstance().Engine().Table(TableName()).In(column, args).Find(&list)
	return list, err
}

// 排除指定字段集合查询
func FindNotIn(column string, args ...interface{}) ([]SysDept, error) {
	var list []SysDept
	err := db.GetInstance().Engine().Table(TableName()).NotIn(column, args).Find(&list)
	return list, err
}

// 批量删除
func DeleteBatch(ids ...int64) (int64, error) {
	return db.GetInstance().Engine().Table(TableName()).In("dept_id", ids).Delete(new(SysDept))
}
