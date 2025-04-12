package model

import (
	"common/models"
	"github.com/lostvip-com/lv_framework/lv_db"
)

// SysMenu 菜单权限
type SysMenu struct {
	MenuId   int64  `gorm:"type:bigint(20);primary_key;auto_increment;菜单ID;" json:"menuId"`
	MenuName string `gorm:"type:varchar(50);comment:菜单名称;" json:"menuName"`
	ParentId int64  `gorm:"type:bigint(20);comment:父菜单ID;" json:"parentId"`
	OrderNum int    `gorm:"type:int(11);comment:显示顺序;" json:"orderNum"`
	Url      string `gorm:"type:varchar(200);comment:请求地址;" json:"url"`
	Target   string `gorm:"type:varchar(20);comment:打开方式（menuItem页签 menuBlank新窗口）;" json:"target"`
	MenuType string `gorm:"type:char(1);comment:菜单类型（M目录 C菜单 F按钮）;" json:"menuType"`
	Visible  string `gorm:"type:char(1);comment:菜单状态（0显示 1隐藏）;" json:"visible"`
	Perms    string `gorm:"type:varchar(100);comment:权限标识;" json:"perms"`
	Icon     string `gorm:"type:varchar(100);comment:菜单图标;" json:"icon"`
	UpdateBy string `gorm:"type:varchar(64);comment:更新者;" json:"updateBy"`
	Remark   string `gorm:"type:varchar(500);comment:备注;" json:"remark"`
	CreateBy string `gorm:"type:varchar(32);comment:创建人;column:create_by;"  json:"createBy"`
	models.BaseModel
	//
	Children   []*SysMenu `gorm:"-" json:"children"`
	ParentName string     `gorm:"-" json:"parentName"`
}

func (e *SysMenu) TableName() string {
	return "sys_menu"
}

// 增
func (e *SysMenu) Insert() error {
	return lv_db.GetMasterGorm().Save(e).Error
}

// 查
func (e *SysMenu) FindById() error {
	err := lv_db.GetMasterGorm().Take(e, e.MenuId).Error
	return err
}

// 查第一条
func (e *SysMenu) FindOne() error {
	tb := lv_db.GetMasterGorm().Table("sys_menu")
	if e.MenuId != 0 {
		tb = tb.Where("menu_id=?", e.MenuId)
	}
	if e.ParentId != 0 {
		tb = tb.Where("parent_id=?", e.ParentId)
	}
	if e.MenuName != "" {
		tb = tb.Where("menu_name=?", e.MenuName)
	}

	if e.Perms != "" {
		tb = tb.Where("perms=?", e.Perms)
	}

	err := tb.First(e).Error
	return err
}

// 查第一条
func (e *SysMenu) FindLastOne() error {
	tb := lv_db.GetMasterGorm().Table("sys_menu")
	if e.MenuId != 0 {
		tb = tb.Where("menu_id=?", e.MenuId)
	}
	if e.ParentId != 0 {
		tb = tb.Where("parent_id=?", e.ParentId)
	}
	if e.MenuName != "" {
		tb = tb.Where("menu_name=?", e.MenuName)
	}
	if e.Perms != "" {
		tb = tb.Where("perms=?", e.Perms)
	}
	tb.Order("menu_id desc")
	tb.Limit(1)
	err := tb.First(e).Error
	return err
}

// 改
func (e *SysMenu) Update() error {
	return lv_db.GetMasterGorm().Table(e.TableName()).Updates(e).Error
}

// 删
func (e *SysMenu) Delete() error {
	return lv_db.GetMasterGorm().Table(e.TableName()).Delete(e).Error
}
