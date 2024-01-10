package menu

import (
	"errors"
	"lostvip.com/db"
	"lostvip.com/utils/lv_web"
)

// SysMenu is the golang structure for table sys_menu.
type EntityExtend struct {
	SysMenu    `xorm:"extends"`
	ParentName string         `json:"parentName"` // 父菜单名称
	Children   []EntityExtend `json:"children"`   // 子菜单
}

// 检查菜单名称请求参数
type CheckMenuNameReq struct {
	MenuId   int64  `form:"menuId"  binding:"required"`
	ParentId int64  `form:"parentId"  binding:"required"`
	MenuName string `form:"menuName"  binding:"required"`
}

// 检查菜单名称请求参数
type CheckMenuNameALLReq struct {
	ParentId int64  `form:"parentId"  binding:"required"`
	MenuName string `form:"menuName"  binding:"required"`
}

// 分页请求参数
type SelectPageReq struct {
	MenuName  string `form:"menuName"`      //菜单名称
	Visible   string `form:"visible"`       //状态
	BeginTime string `form:"beginTime"`     //开始时间
	EndTime   string `form:"endTime"`       //结束时间
	PageNum   int    `form:"pageNum"`       //当前页码
	PageSize  int    `form:"pageSize"`      //每页数
	SortName  string `form:"orderByColumn"` //排序字段
	SortOrder string `form:"isAsc"`         //排序方式
}

// 新增页面请求参数
type AddReq struct {
	ParentId int64  `form:"parentId"  binding:"required"`
	MenuType string `form:"menuType"  binding:"required"`
	MenuName string `form:"menuName"  binding:"required"`
	OrderNum int    `form:"orderNum" binding:"required"`
	Url      string `form:"url"`
	Icon     string `form:"icon"`
	Target   string `form:"target"`
	Perms    string `form:"perms""`
	Visible  string `form:"visible"`
}

// 修改页面请求参数
type EditReq struct {
	MenuId   int64  `form:"menuId" binding:"required"`
	ParentId int64  `form:"parentId"  binding:"required"`
	MenuType string `form:"menuType"  binding:"required"`
	MenuName string `form:"menuName"  binding:"required"`
	OrderNum int    `form:"orderNum" binding:"required"`
	Url      string `form:"url"`
	Icon     string `form:"icon"`
	Target   string `form:"target"`
	Perms    string `form:"perms""`
	Visible  string `form:"visible"`
}

// 根据主键查询数据
func SelectRecordById(id int64) (*EntityExtend, error) {
	db := db.GetInstance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	var result EntityExtend
	model := db.Table(TableName()).Alias("t")
	model.Select("t.menu_id, t.parent_id, t.menu_name, t.order_num, t.url, t.target, t.menu_type, t.visible, t.perms, t.icon, t.remark,(SELECT menu_name FROM sys_menu WHERE menu_id = t.parent_id) parent_name")
	model.Where("menu_id=?", id)
	_, err := model.Get(&result)
	if err != nil {
		return nil, errors.New("获取数据失败")
	}
	return &result, nil
}

// 根据条件分页查询数据
func SelectListPage(param *SelectPageReq) (*[]SysMenu, *lv_web.Paging, error) {
	db := db.GetInstance().Engine()
	p := new(lv_web.Paging)
	if db == nil {
		return nil, p, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName()).Alias("t")

	if param != nil {
		if param.MenuName != "" {
			model.Where("m.menu_name like ?", "%"+param.MenuName+"%")
		}

		if param.Visible != "" {
			model.Where("m.visible = ", param.Visible)
		}

		if param.BeginTime != "" {
			model.Where("date_format(m.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(m.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}

	totalModel := model.Clone()

	total, err := totalModel.Count()

	if err != nil {
		return nil, p, errors.New("读取行数失败")
	}

	p = lv_web.CreatePaging(param.PageNum, param.PageSize, int(total))

	model.Limit(p.Pagesize, p.StartNum)

	var result []SysMenu

	err = model.Find(&result)

	if err != nil {
		return nil, p, errors.New("读取数据失败")
	} else {
		return &result, p, nil
	}
}

// 获取所有数据
func SelectListAll(param *SelectPageReq) ([]SysMenu, error) {
	db := db.GetInstance().Engine()

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName()).Alias("t")

	if param != nil {

		if param.MenuName != "" {
			model.Where("t.menu_name like ?", "%"+param.MenuName+"%")
		}

		if param.Visible != "" {
			model.Where("t.visible = ?", param.Visible)
		}

		if param.BeginTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d') ", param.BeginTime)
		}

		if param.EndTime != "" {
			model.Where("date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d') ", param.EndTime)
		}
	}
	model.OrderBy("t.parent_id,t.order_num")
	var result []SysMenu

	err := model.Find(&result)

	if err != nil {
		return nil, errors.New("读取数据失败")
	} else {
		return result, nil
	}
}

// 获取管理员菜单数据
func SelectMenuNormalAll() ([]EntityExtend, error) {
	var result []EntityExtend

	db := db.GetInstance().Engine()
	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName()).Alias("m")
	model.Select("distinct m.menu_id, m.parent_id, m.menu_name, m.url, m.visible, ifnull(m.perms,'') as perms, m.target, m.menu_type, m.icon, m.order_num, m.create_time")
	model.Where(" m.visible = 0")
	model.OrderBy("m.parent_id, m.order_num ")
	err := model.Find(&result)

	if err != nil {
		return nil, errors.New("读取数据失败")
	} else {
		return result, nil
	}
}

// 根据用户ID读取菜单数据
func SelectMenusByUserId(userId string) ([]EntityExtend, error) {
	var result []EntityExtend

	db := db.GetInstance().Engine()
	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}

	model := db.Table(TableName()).Alias("m")
	model.Join("LEFT", []string{"sys_role_menu", "rm"}, "m.menu_id = rm.menu_id")
	model.Join("LEFT", []string{"sys_user_role", "ur"}, "rm.role_id = ur.role_id")
	model.Join("LEFT", []string{"sys_role", "ro"}, "ur.role_id = ro.role_id")
	model.Select("distinct m.menu_id, m.parent_id, m.menu_name, m.url, m.visible, ifnull(m.perms,'') as perms, m.target, m.menu_type, m.icon, m.order_num, m.create_time")
	model.Where("ur.user_id = ? and  m.visible = 0  AND ro.status = 0", userId)
	model.OrderBy("m.parent_id, m.order_num ")
	err := model.Find(&result)

	if err != nil {
		return nil, errors.New("读取数据失败")
	} else {
		return result, nil
	}
}

// 根据角色ID查询菜单
func SelectMenuTree(roleId int64) ([]string, error) {
	db := db.GetInstance().Engine()

	var result []string

	if db == nil {
		return nil, errors.New("获取数据库连接失败")
	}
	model := db.Table(TableName()).Alias("m")
	model.Join("LEFT", []string{"sys_role_menu", "rm"}, "m.menu_id = rm.menu_id")
	model.Where("rm.role_id = ?", roleId)
	model.OrderBy("m.parent_id, m.order_num ")
	model.Select("concat(m.menu_id, ifnull(m.perms,'')) as perms")
	var list []SysMenu
	err := model.Find(&list)
	if err != nil {
		return nil, errors.New("读取数据失败")
	}

	for _, record := range list {
		if record.Perms != "" {
			result = append(result, record.Perms)
		}
	}

	return result, nil
}

// 校验菜单名称是否唯一
func CheckMenuNameUniqueAll(menuName string, parentId int64) (*SysMenu, error) {
	var entity SysMenu
	entity.MenuName = menuName
	entity.ParentId = parentId
	ok, err := entity.FindOne()
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}

// 校验菜单名称是否唯一
func CheckPermsUniqueAll(perms string) (*SysMenu, error) {
	var entity SysMenu
	entity.Perms = perms
	ok, err := entity.FindOne()
	if ok {
		return &entity, err
	} else {
		return nil, err
	}
}
