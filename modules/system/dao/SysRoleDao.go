package dao

import (
	"common/common_vo"
	"errors"
	"github.com/lostvip-com/lv_framework/lv_db"
	"system/model"
)

type SysRoleDao struct {
}

// 根据条件分页查询角色数据
func (dao *SysRoleDao) SelectListPage(param *common_vo.RolePageReq) (result []model.SysRole, total int64, err error) {
	db := lv_db.GetMasterGorm()
	if db == nil {
		return nil, 0, errors.New("获取数据库连接失败")
	}
	tb := db.Table("sys_role as r").Where("r.del_flag = '0'")
	if param.RoleName != "" {
		tb.Where("r.role_name like ?", "%"+param.RoleName+"%")
	}
	if param.Status != "" {
		tb.Where("r.status = ?", param.Status)
	}

	if param.RoleKey != "" {
		tb.Where("r.role_key like ?", "%"+param.RoleKey+"%")
	}

	if param.DataScope != "" {
		tb.Where("r.data_scope = ?", param.DataScope)
	}

	if param.BeginTime != "" {
		tb.Where("date_format(r.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
	}

	if param.EndTime != "" {
		tb.Where("date_format(r.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
	}

	if param.SortName != "" {
		tb.Order(param.SortName + " " + param.SortOrder + " ")
	}
	err = tb.Count(&total).Offset(param.GetStartNum()).Limit(param.GetPageSize()).Find(&result).Error
	return result, total, err
}

// 获取所有角色数据
func (dao *SysRoleDao) SelectListAll(param *common_vo.RolePageReq) ([]common_vo.SysRoleFlag, error) {
	db := lv_db.GetMasterGorm()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table("sys_role as r").Select("r.*,false as flag").Where("r.del_flag = '0'")
	if param != nil {
		if param.RoleName != "" {
			model.Where("r.role_name like ?", "%"+param.RoleName+"%")
		}

		if param.Status != "" {
			model.Where("r.status = ", param.Status)
		}

		if param.RoleKey != "" {
			model.Where("r.role_key like ?", "%"+param.RoleKey+"%")
		}

		if param.DataScope != "" {
			model.Where("r.data_scope = ", param.DataScope)
		}

		if param.BeginTime != "" {
			model.Where("date_format(r.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(r.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	var result []common_vo.SysRoleFlag

	err := model.Find(&result).Error
	return result, err
}

// 根据用户ID查询角色
func (dao *SysRoleDao) SelectRoleContactVo(userId int64) ([]model.SysRole, error) {
	db := lv_db.GetMasterGorm()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	tb := db.Table("sys_role as r")
	tb.Joins("Left join sys_user_role as ur on ur.role_id = r.role_id")
	tb.Joins("Left join sys_user as u on u.user_id = ur.user_id")
	tb.Joins("Left join sys_dept as d on u.dept_id = d.dept_id")
	tb.Where("r.del_flag = '0'")
	tb.Where("ur.user_id = ?", userId)
	tb.Select("distinct r.role_id, r.role_name, r.role_key, r.role_sort, r.data_scope,r.status, r.del_flag, r.create_time, r.remark")

	var result []model.SysRole

	err := tb.Find(&result).Error
	return result, err
}

// 检查角色键是否唯一
func (dao *SysRoleDao) FindRoleByName(roleName string) (*model.SysRole, error) {
	var entity model.SysRole
	entity.RoleName = roleName
	err := entity.FindOne()
	return &entity, err
}

// 检查角色键是否唯一
func (dao *SysRoleDao) FindRoleByRoleKey(roleKey string) (*model.SysRole, error) {
	var entity model.SysRole
	entity.RoleKey = roleKey
	err := entity.FindOne()
	return &entity, err
}

func (e *SysRoleDao) FindCount(roleKey, roleName string) (int64, error) {
	var count int64 = 0
	tb := lv_db.GetMasterGorm()
	if roleName != "" {
		tb = tb.Where("role_name=? and del_flag=0", roleName)
	}
	if roleKey != "" {
		tb = tb.Where("role_key=? and del_flag=0", roleKey)
	}
	err := tb.Count(&count).Error
	return count, err
}
