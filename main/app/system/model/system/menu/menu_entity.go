package menu

import (
	"lostvip.com/db"
	"time"
)

type SysMenu struct {
	MenuId     int64     `json:"menu_id" xorm:"not null pk autoincr comment('菜单ID') BIGINT(20)"`
	MenuName   string    `json:"menu_name" xorm:"not null comment('菜单名称') VARCHAR(50)"`
	ParentId   int64     `json:"parent_id" xorm:"default 0 comment('父菜单ID') BIGINT(20)"`
	OrderNum   int       `json:"order_num" xorm:"default 0 comment('显示顺序') INT(4)"`
	Url        string    `json:"url" xorm:"default '#' comment('请求地址') VARCHAR(200)"`
	Target     string    `json:"target" xorm:"default '' comment('打开方式（menuItem页签 menuBlank新窗口）') VARCHAR(20)"`
	MenuType   string    `json:"menu_type" xorm:"default '' comment('菜单类型（M目录 C菜单 F按钮）') CHAR(1)"`
	Visible    string    `json:"visible" xorm:"default '0' comment('菜单状态（0显示 1隐藏）') CHAR(1)"`
	Perms      string    `json:"perms" xorm:"comment('权限标识') VARCHAR(100)"`
	Icon       string    `json:"icon" xorm:"default '#' comment('菜单图标') VARCHAR(100)"`
	CreateBy   string    `json:"create_by" xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `json:"update_by" xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('更新时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"default '' comment('备注') VARCHAR(500)"`
}

// 映射数据表
func TableName() string {
	return "sys_menu"
}

// 插入数据
func (r *SysMenu) Insert() (int64, error) {
	return db.Instance().Engine().Table(TableName()).Insert(r)
}

// 更新数据
func (r *SysMenu) Update() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.MenuId).Update(r)
}

// 删除
func (r *SysMenu) Delete() (int64, error) {
	return db.Instance().Engine().Table(TableName()).ID(r.MenuId).Delete(r)
}

// 批量删除
func DeleteBatch(ids ...int64) (int64, error) {
	return db.Instance().Engine().Table(TableName()).In("menu_id", ids).Delete(new(SysMenu))
}

func DeleteChildren(parentId int64) (int64, error) {
	return db.Instance().Engine().Table(TableName()).Where("parent_id", parentId).Delete(new(SysMenu))
}

// 根据结构体中已有的非空数据来获得单条数据
func (r *SysMenu) FindOne() (bool, error) {
	return db.Instance().Engine().Table(TableName()).OrderBy("menu_id desc").Get(r)
}

// 根据条件查询
func Find(where, order string) ([]SysMenu, error) {
	var list []SysMenu
	err := db.Instance().Engine().Table(TableName()).Where(where).OrderBy(order).Find(&list)
	return list, err
}

// 指定字段集合查询
func FindIn(column string, args ...interface{}) ([]SysMenu, error) {
	var list []SysMenu
	err := db.Instance().Engine().Table(TableName()).In(column, args).Find(&list)
	return list, err
}

// 排除指定字段集合查询
func FindNotIn(column string, args ...interface{}) ([]SysMenu, error) {
	var list []SysMenu
	err := db.Instance().Engine().Table(TableName()).NotIn(column, args).Find(&list)
	return list, err
}
