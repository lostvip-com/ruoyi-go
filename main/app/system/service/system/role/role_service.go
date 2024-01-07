package role

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lostvip.com/db"
	"lostvip.com/utils/lv_conv"
	"lostvip.com/utils/lv_web"
	"robvi/app/system/model/system/role"
	"robvi/app/system/model/system/role_dept"
	"robvi/app/system/model/system/role_menu"
	"robvi/app/system/model/system/user_role"
	"robvi/app/system/service"
	"time"
)

// 根据主键查询数据
func SelectRecordById(id int64) (*role.SysRole, error) {
	entity := &role.SysRole{RoleId: id}
	_, err := entity.FindOne()
	return entity, err
}

// 根据条件查询数据
func SelectRecordAll(params *role.SelectPageReq) ([]role.EntityFlag, error) {
	return role.SelectListAll(params)
}

// 根据条件分页查询数据
func SelectRecordPage(params *role.SelectPageReq) ([]role.SysRole, *lv_web.Paging, error) {
	return role.SelectListPage(params)
}

// 根据主键删除数据
func DeleteRecordById(id int64) bool {
	rs, _ := (&role.SysRole{RoleId: id}).Delete()
	if rs > 0 {
		return true
	}
	return false
}

// 添加数据
func AddSave(req *role.AddReq, c *gin.Context) (int64, error) {

	role := new(role.SysRole)
	role.RoleName = req.RoleName
	role.RoleKey = req.RoleKey
	role.Status = req.Status
	role.Remark = req.Remark
	role.CreateTime = time.Now()
	role.CreateBy = ""
	role.DelFlag = "0"
	role.DataScope = "1"
	var userService service.UserService
	user := userService.GetProfile(c)

	if user != nil {
		role.CreateBy = user.LoginName
	}

	session := db.GetInstance().Engine().NewSession()

	err := session.Begin()

	_, err = session.Insert(role)

	if err != nil {
		session.Rollback()
		return 0, err
	}

	if req.MenuIds != "" {
		menus := lv_conv.ToInt64Array(req.MenuIds, ",")
		if len(menus) > 0 {
			roleMenus := make([]role_menu.Entity, 0)
			for i := range menus {
				if menus[i] > 0 {
					var tmp role_menu.Entity
					tmp.RoleId = role.RoleId
					tmp.MenuId = menus[i]
					roleMenus = append(roleMenus, tmp)
				}
			}
			if len(roleMenus) > 0 {
				_, err := session.Table(role_menu.TableName()).Insert(roleMenus)
				if err != nil {
					session.Rollback()
					return 0, err
				}
			}
		}
	}
	err = session.Commit()
	return role.RoleId, err
}

// 修改数据
func EditSave(req *role.EditReq, c *gin.Context) (int64, error) {
	r := &role.SysRole{RoleId: req.RoleId}
	ok, err := r.FindOne()
	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("数据不存在")
	}
	r.RoleName = req.RoleName
	r.RoleKey = req.RoleKey
	r.Status = req.Status
	r.Remark = req.Remark
	r.UpdateTime = time.Now()
	r.UpdateBy = ""
	var userService service.UserService
	user := userService.GetProfile(c)

	if user == nil {
		r.CreateBy = user.LoginName
	}

	session := db.GetInstance().Engine().NewSession()

	pErr := session.Begin()

	_, pErr = session.Table("sys_role").ID(r.RoleId).Update(r)

	if pErr != nil {
		session.Rollback()
		return 0, pErr
	}

	if req.MenuIds != "" {
		menus := lv_conv.ToInt64Array(req.MenuIds, ",")
		if len(menus) > 0 {
			roleMenus := make([]role_menu.Entity, 0)
			for i := range menus {
				if menus[i] > 0 {
					var tmp role_menu.Entity
					tmp.RoleId = r.RoleId
					tmp.MenuId = menus[i]
					roleMenus = append(roleMenus, tmp)
				}
			}
			if len(roleMenus) > 0 {
				_, pErr = session.Exec("delete from sys_role_menu where role_id=?", r.RoleId)
				if pErr != nil {
					session.Rollback()
					return 0, pErr
				}
				_, pErr = session.Table(role_menu.TableName()).Insert(roleMenus)
				if pErr != nil {
					session.Rollback()
					return 0, err
				}
			}
		}
	}
	return 1, session.Commit()
}

// 保存数据权限
func AuthDataScope(req *role.DataScopeReq, c *gin.Context) (int64, error) {
	entity := &role.SysRole{RoleId: req.RoleId}
	ok, err := entity.FindOne()
	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("数据不存在")
	}

	if req.DataScope != "" {
		entity.DataScope = req.DataScope
	}
	var userService service.UserService
	user := userService.GetProfile(c)

	if user != nil {
		entity.UpdateBy = user.LoginName
	}
	entity.UpdateTime = time.Now()

	session := db.GetInstance().Engine().NewSession()
	tanErr := session.Begin()

	_, tanErr = session.Table("sys_role").ID(entity.RoleId).Update(entity)

	if tanErr != nil {
		session.Rollback()
		return 0, err
	}

	if req.DeptIds != "" {
		deptids := lv_conv.ToInt64Array(req.DeptIds, ",")
		if len(deptids) > 0 {
			roleDepts := make([]role_dept.Entity, 0)
			for i := range deptids {
				if deptids[i] > 0 {
					var tmp role_dept.Entity
					tmp.RoleId = entity.RoleId
					tmp.DeptId = deptids[i]
					roleDepts = append(roleDepts, tmp)
				}
			}
			if len(roleDepts) > 0 {
				session.Exec("delete from  sys_role_dept where role_id=?", entity.RoleId)
				_, tanErr := session.Table(role_dept.TableName()).Insert(roleDepts)
				if tanErr != nil {
					session.Rollback()
					return 0, err
				}
			}
		}
	}
	return 1, session.Commit()

}

// 批量删除数据记录
func DeleteRecordByIds(ids string) (int64, error) {
	idArr := lv_conv.ToInt64Array(ids, ",")
	result, _ := db.GetInstance().Engine().Exec("delete from sys_role where role_id in ? ", idArr)
	return result.RowsAffected()
}

// 导出excel
func Export(param *role.SelectPageReq) (string, error) {
	head := []string{"角色ID", "角色名称", "权限字符串", "显示顺序", "数据范围", "角色状态"}
	col := []string{"role_id", "role_name", "role_key", "role_sort", "data_scope", "status"}
	return role.SelectListExport(param, head, col)
}

// 根据用户ID查询角色
func SelectRoleContactVo(userId int64) ([]role.EntityFlag, error) {
	var paramsPost *role.SelectPageReq
	roleAll, err := role.SelectListAll(paramsPost)
	if err != nil || roleAll == nil {
		return nil, errors.New("未查询到角色数据")
	}

	userRole, err := role.SelectRoleContactVo(userId)
	if userRole != nil {
		for i := range userRole {
			for j := range roleAll {
				if userRole[i].RoleId == roleAll[j].RoleId {
					roleAll[j].Flag = true
					break
				}
			}
		}
	}
	return roleAll, nil
}

// 批量选择用户授权
func InsertAuthUsers(roleId int64, userIds string) int64 {
	idarr := lv_conv.ToInt64Array(userIds, ",")
	var roleUserList []user_role.SysUserRole
	for _, str := range idarr {
		var tmp user_role.SysUserRole
		tmp.UserId = str
		tmp.RoleId = roleId
		roleUserList = append(roleUserList, tmp)
	}

	rs, err := db.GetInstance().Engine().Table("sys_user_role").Insert(roleUserList)
	if err != nil {
		return 0
	}
	return rs
}

// 取消授权用户角色
func DeleteUserRoleInfo(userId, roleId int64) int64 {
	entity := &user_role.SysUserRole{UserId: userId, RoleId: roleId}
	rs, _ := entity.Delete()
	return rs
}

// 批量取消授权用户角色
func DeleteUserRoleInfos(roleId int64, ids string) int64 {
	idarr := lv_conv.ToInt64Array(ids, ",")

	idStr := ""
	for _, item := range idarr {
		if item > 0 {
			if idStr != "" {
				idStr += "," + lv_conv.String(item)
			} else {
				idStr = lv_conv.String(item)
			}
		}
	}

	rs, err := db.GetInstance().Engine().Exec("delete from sys_user_role where role_id=? and user_id in (?)", roleId, idStr)
	if err != nil {
		return 0
	}
	nums, err := rs.RowsAffected()
	if err != nil {
		return 0
	}
	return nums
}

// 检查角色名是否唯一
func CheckRoleNameUniqueAll(roleName string) string {
	entity, err := role.CheckRoleNameUniqueAll(roleName)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.RoleId > 0 {
		return "1"
	}
	return "0"
}

// 检查角色键是否唯一
func CheckRoleKeyUniqueAll(roleKey string) string {
	entity, err := role.CheckRoleKeyUniqueAll(roleKey)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.RoleId > 0 {
		return "1"
	}
	return "0"
}

// 检查角色名是否唯一
func CheckRoleNameUnique(roleName string, roleId int64) string {
	entity, err := role.CheckRoleNameUniqueAll(roleName)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.RoleId > 0 && entity.RoleId != roleId {
		return "1"
	}
	return "0"
}

// 检查角色键是否唯一
func CheckRoleKeyUnique(roleKey string, roleId int64) string {
	entity, err := role.CheckRoleKeyUniqueAll(roleKey)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.RoleId > 0 && entity.RoleId != roleId {
		return "1"
	}
	return "0"
}

// 判断是否是管理员
func IsAdmin(id int64) bool {
	if id == 1 {
		return true
	} else {
		return false
	}
}

// 校验角色是否允许操作
func CheckRoleAllowed(id int64) bool {
	if IsAdmin(id) {
		return false
	} else {
		return true
	}
}
