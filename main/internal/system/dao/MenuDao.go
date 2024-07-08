package dao

import (
	"errors"
	"github.com/lv_framework/db"
	"github.com/lv_framework/utils/lv_err"
	"main/internal/system/model"
	"main/internal/system/vo"
)

// 修改页面请求参数
type MenuDao struct {
}

// 批量删除
func (r *MenuDao) DeleteBatch(ids ...int64) (int64, error) {
	db := db.GetMasterGorm().Table("sys_menu").Where("menu_id in ? ", ids).Update("del_flag", 1)
	return db.RowsAffected, db.Error
}

func (r *MenuDao) DeleteChildren(parentId int64) (int64, error) {
	tb := db.GetMasterGorm().Table("sys_menu").Where("parent_id=?", parentId).Update("del_flag", 1)
	return tb.RowsAffected, tb.Error
}

// 根据主键查询数据
func (dao *MenuDao) SelectRecordById(id int64) (*model.SysMenu, error) {
	tb := db.GetMasterGorm()
	if tb == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	var result model.SysMenu
	tb = tb.Table("sys_menu t")
	tb.Select("t.menu_id, t.parent_id, t.menu_name, t.order_num, t.url, t.target, t.menu_type, t.visible, t.perms, t.icon, t.remark,(SELECT menu_name FROM sys_menu WHERE menu_id = t.parent_id) parent_name")
	tb.Where("t.menu_id=?", id)
	err := tb.First(&result).Error
	if err != nil {
		return nil, errors.New("获取数据失败")
	}
	return &result, nil
}

// 根据条件分页查询数据
func (dao *MenuDao) SelectListPage(param *vo.SelectMenuPageReq) (*[]model.SysMenu, int64, error) {
	tb := db.GetMasterGorm()
	tb = tb.Table("sys_menu")
	if param != nil {
		if param.MenuName != "" {
			tb.Where("menu_name like ?", "%"+param.MenuName+"%")
		}

		if param.Visible != "" {
			tb.Where("visible = ", param.Visible)
		}

		if param.BeginTime != "" {
			tb.Where("date_format(create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			tb.Where("date_format(create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	var result []model.SysMenu
	err := tb.Offset(param.GetStartNum()).Limit(param.GetPageSize()).Find(&result).Error
	lv_err.HasErrAndPanic(err)
	err = tb.Offset(0).Limit(-1).Count(&param.Total).Error
	lv_err.HasErrAndPanic(err)
	return &result, param.Total, nil
}

// 获取所有数据
func (dao *MenuDao) SelectListAll(param *vo.SelectMenuPageReq) ([]model.SysMenu, error) {
	tb := db.GetMasterGorm()
	tb = tb.Table("sys_menu as t")

	if param != nil {

		if param.MenuName != "" {
			tb.Where("t.menu_name like ?", "%"+param.MenuName+"%")
		}

		if param.Visible != "" {
			tb.Where("t.visible = ?", param.Visible)
		}

		if param.BeginTime != "" {
			tb.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			tb.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	tb.Order("t.parent_id,t.order_num desc")
	var result []model.SysMenu

	err := tb.Find(&result).Error

	if err != nil {
		return nil, err
	} else {
		return result, err
	}
}

// 获取管理员菜单数据
func (dao *MenuDao) SelectMenuNormalAll() ([]model.SysMenu, error) {
	var result []model.SysMenu

	tb := db.GetMasterGorm()
	tb = tb.Table("sys_menu as m")
	tb.Select("distinct m.menu_id, m.parent_id, m.menu_name, m.url, m.visible, ifnull(m.perms,'') as perms, m.target, m.menu_type, m.icon, m.order_num, m.create_time")
	tb.Where(" m.visible = 0")
	tb.Order("m.parent_id, m.order_num")
	err := tb.Find(&result).Error

	if err != nil {
		return nil, err
	} else {
		return result, err
	}
}

// 根据用户ID读取菜单数据
func (dao *MenuDao) SelectMenusByUserId(userId int64) ([]model.SysMenu, error) {
	var result []model.SysMenu

	db := db.GetMasterGorm()
	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	tb := db.Table("sys_menu as m")
	tb.Joins("LEFT join sys_role_menu as rm on m.menu_id = rm.menu_id")
	tb.Joins("LEFT join sys_user_role as ur on rm.role_id = ur.role_id")
	tb.Joins("LEFT join sys_role as ro on ur.role_id = ro.role_id")
	tb.Select("distinct m.menu_id, m.parent_id, m.menu_name, m.url, m.visible, ifnull(m.perms,'') as perms, m.target, m.menu_type, m.icon, m.order_num, m.create_time")
	tb.Where("ur.user_id = ? and  m.visible = 0  AND ro.status = 0", userId)
	tb.Order("m.parent_id, m.order_num")
	err := tb.Find(&result).Error

	if err != nil {
		return nil, err
	} else {
		return result, err
	}
}

// 根据角色ID查询菜单
func (dao *MenuDao) SelectMenuTree(roleId int64) ([]string, error) {
	tb := db.GetMasterGorm()
	if tb == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	var result []string

	if tb == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	tb = tb.Table("sys_menu as m ")
	tb.Joins(" LEFT join sys_role_menu as rm on m.menu_id = rm.menu_id")
	tb.Where(" rm.role_id = ?", roleId)
	tb.Order(" m.parent_id, m.order_num desc ")
	tb.Select("concat(m.menu_id, ifnull(m.perms,'')) as perms")
	var list []model.SysMenu
	err := tb.Find(&list).Error
	if err != nil {
		return nil, err
	}
	for _, record := range list {
		if record.Perms != "" {
			result = append(result, record.Perms)
		}
	}

	return result, err
}

// 校验菜单名称是否唯一
func (dao *MenuDao) CheckMenuNameExists(menuName string, parentId int64) bool {
	var entity model.SysMenu
	entity.MenuName = menuName
	entity.ParentId = parentId
	err := entity.FindOne()
	if err == nil && entity.MenuId > 0 {
		return true
	} else {
		return false
	}
}

// 校验菜单名称是否唯一
func (dao *MenuDao) CheckPermsUniqueAll(perms string) (*model.SysMenu, error) {
	var entity model.SysMenu
	entity.Perms = perms
	err := entity.FindOne()
	return &entity, err
}
