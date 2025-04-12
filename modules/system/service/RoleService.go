package service

import (
	"common/common_vo"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_db"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"gorm.io/gorm"
	"system/dao"
	"system/model"
	"time"
)

type RoleService struct{}

// 根据主键查询数据
func (svc *RoleService) SelectRecordById(id int64) (*model.SysRole, error) {
	entity := &model.SysRole{RoleId: id}
	err := entity.FindOne()
	return entity, err
}

// 根据主键查询数据
func (svc *RoleService) SelectRecordPage(params *common_vo.RolePageReq) ([]model.SysRole, int64, error) {
	var d dao.SysRoleDao
	return d.SelectListPage(params)
}

// 根据条件查询数据
func (svc *RoleService) SelectRecordAll(params *common_vo.RolePageReq) ([]common_vo.SysRoleFlag, error) {
	var dao dao.SysRoleDao
	return dao.SelectListAll(params)
}

// 添加数据
func (svc *RoleService) AddSave(req *common_vo.AddRoleReq, c *gin.Context) (int64, error) {

	role := new(model.SysRole)
	role.RoleName = req.RoleName
	role.RoleKey = req.RoleKey
	role.Status = req.Status
	role.Remark = req.Remark
	role.CreateTime = time.Now()
	role.DataScope = "1"
	var userService UserService
	user := userService.GetProfile(c)

	if user != nil {
		role.CreateBy = user.LoginName
	}
	session := lv_db.GetMasterGorm().Begin()
	err := session.Save(role).Error
	lv_err.HasErrAndPanic(err)
	if err != nil {
		session.Rollback()
		return 0, err
	}

	if req.MenuIds != "" {
		menus := lv_conv.ToInt64Array(req.MenuIds, ",")
		if len(menus) > 0 {
			roleMenus := make([]model.SysRoleMenu, 0)
			for i := range menus {
				if menus[i] > 0 {
					var tmp model.SysRoleMenu
					tmp.RoleId = role.RoleId
					tmp.MenuId = menus[i]
					roleMenus = append(roleMenus, tmp)
				}
			}
			if len(roleMenus) > 0 {
				err := session.Save(roleMenus).Error
				if err != nil {
					session.Rollback()
					return 0, err
				}
			}
		}
	}
	err = session.Commit().Error
	return role.RoleId, err
}

// 修改数据
func (svc *RoleService) EditSave(req *common_vo.EditRoleReq, c *gin.Context) (int64, error) {
	r := &model.SysRole{RoleId: req.RoleId}
	err := r.FindOne()
	lv_err.HasErrAndPanic(err)
	r.RoleName = req.RoleName
	r.RoleKey = req.RoleKey
	r.Status = req.Status
	r.Remark = req.Remark
	r.UpdateTime = time.Now()
	r.UpdateBy = ""
	var userService UserService
	user := userService.GetProfile(c)
	if user == nil {
		r.CreateBy = user.LoginName
	}
	db := lv_db.GetMasterGorm()
	err = db.Transaction(func(tx *gorm.DB) error {
		//更新role表
		if err := tx.Updates(&r).Error; err != nil {
			return err
		}
		//删除旧的功能权限授权，中间表
		if req.MenuIds != "" {
			menus := lv_conv.ToInt64Array(req.MenuIds, ",")
			if len(menus) > 0 {
				roleMenus := make([]model.SysRoleMenu, 0)
				for i := range menus {
					if menus[i] > 0 {
						var tmp model.SysRoleMenu
						tmp.RoleId = r.RoleId
						tmp.MenuId = menus[i]
						roleMenus = append(roleMenus, tmp)
					}
				}
				if len(roleMenus) > 0 {
					err = tx.Exec("delete from sys_role_menu where role_id=?", r.RoleId).Error
					if err != nil {
						return err
					}
					//插入新的功能权限授权，中间表
					err = tx.CreateInBatches(roleMenus, len(roleMenus)).Error
					if err != nil {
						return err
					}
				}
			}
		}
		return err
	})

	return 1, err
}

// 保存数据权限
func (svc *RoleService) AuthDataScope(req *common_vo.DataScopeReq, c *gin.Context) (int64, error) {
	entity := &model.SysRole{RoleId: req.RoleId}
	err := entity.FindOne()
	lv_err.HasErrAndPanic(err)
	if req.DataScope != "" {
		entity.DataScope = req.DataScope
	}
	var userService UserService
	user := userService.GetProfile(c)
	if user != nil {
		entity.UpdateBy = user.LoginName
	}
	entity.UpdateTime = time.Now()

	db := lv_db.GetMasterGorm()
	err = db.Transaction(func(tx *gorm.DB) error {
		//更新role表
		if err := tx.Updates(&entity).Error; err != nil {
			return err
		}
		//删除旧的功能权限授权，中间表
		if req.DeptIds != "" {
			deptids := lv_conv.ToInt64Array(req.DeptIds, ",")
			if len(deptids) > 0 {
				roleDepts := make([]model.SysRoleDept, 0)
				for i := range deptids {
					if deptids[i] > 0 {
						var tmp model.SysRoleDept
						tmp.RoleId = entity.RoleId
						tmp.DeptId = deptids[i]
						roleDepts = append(roleDepts, tmp)
					}
				}
				if len(roleDepts) > 0 {
					tx.Exec("delete from  sys_role_dept where role_id=?", entity.RoleId)
					err := tx.CreateInBatches(roleDepts, len(roleDepts)).Error
					return err
				}
			}
		}
		return err
	})

	return 1, err

}

// 批量删除数据记录
func (svc *RoleService) DeleteRecordByIds(ids string) error {
	idArr := lv_conv.ToInt64Array(ids, ",")
	idsDel := make([]int64, 0)
	for _, id := range idArr {
		if id != 1 { //忽略admin
			idsDel = append(idsDel, id)
		}
	}
	err := lv_db.GetMasterGorm().Exec("delete from sys_role where role_id in ?  ", idsDel).Error
	return err
}

// 根据用户ID查询角色
func (svc *RoleService) SelectRoleContactVo(userId int64) ([]common_vo.SysRoleFlag, error) {
	var paramsPost *common_vo.RolePageReq
	var dao dao.SysRoleDao
	roleAll, err := dao.SelectListAll(paramsPost)
	if err != nil || roleAll == nil {
		return nil, errors.New("未查询到角色数据")
	}
	userRole, err := dao.SelectRoleContactVo(userId)
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
func (svc *RoleService) InsertAuthUsers(roleId int64, userIds string) error {
	idarr := lv_conv.ToInt64Array(userIds, ",")
	var roleUserList []model.SysUserRole
	for _, str := range idarr {
		var tmp model.SysUserRole
		tmp.UserId = str
		tmp.RoleId = roleId
		roleUserList = append(roleUserList, tmp)
	}
	err := lv_db.GetMasterGorm().CreateInBatches(roleUserList, len(roleUserList)).Error
	return err
}

// 取消授权用户角色
func (svc *RoleService) DeleteUserRoleInfo(userId, roleId int64) error {
	entity := &model.SysUserRole{UserId: userId, RoleId: roleId}
	if entity.RoleId == 1 {
		panic("禁止取消超级管理员权限")
	}
	err := entity.Delete()
	return err
}

// 批量取消授权用户角色
func (svc *RoleService) DeleteUserRoleInfos(roleId int64, ids string) error {
	idarr := lv_conv.ToInt64Array(ids, ",")

	idStr := ""
	for _, item := range idarr {
		if item == 1 {
			continue
		}
		if item > 0 {
			if idStr != "" {
				idStr += "," + lv_conv.String(item)
			} else {
				idStr = lv_conv.String(item)
			}
		}
	}
	err := lv_db.GetMasterGorm().Exec("delete from sys_user_role where role_id=? and user_id in (?)", roleId, idStr).Error
	return err
}

// 检查角色名是否唯一
func (svc *RoleService) IsRoleNameExist(roleName string) (bool, error) {
	var dao dao.SysRoleDao
	_, err := dao.FindRoleByName(roleName)
	if err == nil {
		return true, err
	}
	return false, err
}

// 检查角色键是否唯一
func (svc *RoleService) IsRoleKeyExist(roleKey string) (bool, error) {
	var dao dao.SysRoleDao
	_, err := dao.FindRoleByRoleKey(roleKey)
	if err == nil {
		return true, err
	}
	return false, err
}

// 判断是否是管理员
func (svc *RoleService) IsAdmin(id int64) bool {
	if id == 1 {
		return true
	} else {
		return false
	}
}

// 校验角色是否允许操作
func (svc *RoleService) CheckRoleAllowed(id int64) bool {
	if svc.IsAdmin(id) {
		return false
	} else {
		return true
	}
}
