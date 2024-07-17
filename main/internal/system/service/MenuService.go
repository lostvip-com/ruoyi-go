package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/db"
	"github.com/lostvip-com/lv_framework/db/lvdao"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_web"
	"github.com/lostvip-com/lv_framework/web/dto"
	"main/internal/system/dao"
	"main/internal/system/model"
	"main/internal/system/vo"
	"strings"
	"time"
)

type MenuService struct {
}

// 根据主键查询数据
func (svc *MenuService) SelectRecordById(id int64) (*model.SysMenu, error) {
	return svc.SelectRecordById(id)
}

// 根据条件查询数据
func (svc *MenuService) SelectListAll(params *vo.SelectMenuPageReq) ([]model.SysMenu, error) {
	return svc.SelectListAll(params)
}

// 根据条件分页查询数据
func (svc *MenuService) SelectListPage(params *vo.SelectMenuPageReq) (*[]model.SysMenu, *lv_web.Paging, error) {
	return svc.SelectListPage(params)
}

// 根据主键删除数据
func (svc *MenuService) DeleteRecordById(id int64) bool {
	err := (&model.SysMenu{MenuId: id}).Delete()
	if err == nil {
		db.GetInstance().Engine().Exec("delete from sys_menu where parent_id=?", id)
		return true
	}
	return false
}

// 添加数据
func (svc *MenuService) AddSave(req *vo.AddMenuReq, c *gin.Context) (int64, error) {

	var entity model.SysMenu
	entity.MenuName = req.MenuName
	entity.Visible = req.Visible
	entity.ParentId = req.ParentId
	entity.Remark = ""
	entity.MenuType = req.MenuType
	entity.Url = req.Url
	entity.Perms = req.Perms
	entity.Target = req.Target
	entity.Icon = req.Icon
	entity.OrderNum = req.OrderNum
	entity.CreateTime = time.Now()
	entity.CreateBy = ""

	var userService UserService
	user := userService.GetProfile(c)

	if user == nil {
		entity.CreateBy = user.LoginName
	}

	err := entity.Insert()
	return entity.MenuId, err
}

// 修改数据
func (svc *MenuService) EditSave(req *vo.EditMenuReq, c *gin.Context) (int64, error) {
	entity := &model.SysMenu{MenuId: req.MenuId}
	err := entity.FindOne()

	if err != nil {
		return 0, err
	}
	entity.MenuName = req.MenuName
	entity.Visible = req.Visible
	entity.ParentId = req.ParentId
	entity.Remark = ""
	entity.MenuType = req.MenuType
	entity.Url = req.Url
	entity.Perms = req.Perms
	entity.Target = req.Target
	entity.Icon = req.Icon
	entity.OrderNum = req.OrderNum
	entity.UpdateTime = time.Now()
	entity.UpdateBy = ""

	var userService UserService
	user := userService.GetProfile(c)

	if user == nil {
		entity.UpdateBy = user.LoginName
	}

	return entity.MenuId, entity.Update()
}

// 批量删除数据记录
func (svc *MenuService) DeleteRecordByIds(ids string) int64 {
	idarr := lv_conv.ToInt64Array(ids, ",")
	var dao dao.MenuDao
	result, err := dao.DeleteBatch(idarr...)
	if err != nil {
		return 0
	}
	return result
}

// 加载所有菜单列表树
func (svc *MenuService) MenuTreeData(userId int64) (*[]dto.Ztree, error) {
	var result *[]dto.Ztree
	menuList, err := svc.SelectMenuNormalByUser(userId)
	if err != nil {
		return nil, err
	}
	result, err = svc.InitZtree(menuList, nil, false)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 获取用户的菜单数据
func (svc *MenuService) SelectMenuNormalByUser(userId int64) (*[]model.SysMenu, error) {
	var userService UserService
	if userService.IsAdmin(userId) {
		return svc.SelectMenuNormalAll()
	} else {
		return svc.SelectMenusByUserId(userId)
	}
}

// 获取用户的菜单数据
func (svc *MenuService) ListMenuNormalByUser(userId int64) (*[]model.SysMenu, error) {
	var dao dao.MenuDao
	var userService UserService
	if userService.IsAdmin(userId) {
		menus, err := dao.SelectMenuNormalAll()
		return &menus, err
	} else {
		menus, err := dao.SelectMenusByUserId(userId)
		return &menus, err
	}
}

// 获取管理员菜单数据
func (svc *MenuService) SelectMenuNormalAll() (*[]model.SysMenu, error) {
	//从数据库中读取
	var dao dao.MenuDao
	menus, err := dao.SelectMenuNormalAll()
	lv_err.HasErrAndPanic(err)
	arr := make([]model.SysMenu, 0)
	childrenMap := svc.InitChildMap(menus)
	for _, menu := range menus {
		if menu.ParentId == 0 { //发现子菜单
			svc.FillAllChildren(&menu, childrenMap)
			arr = append(arr, menu)
		}
	}
	//存入缓存
	return &arr, nil
}

// 根据用户ID读取菜单数据
func (svc *MenuService) SelectMenusByUserId(userId int64) (*[]model.SysMenu, error) {
	//从数据库中读取
	var dao dao.MenuDao
	menus, err := dao.SelectMenusByUserId(userId)
	//存入缓存
	arr := make([]model.SysMenu, 0)
	if err == nil {
		childrenMap := svc.InitChildMap(menus)
		for _, menu := range menus {
			if menu.ParentId == 0 { //发现子菜单
				svc.FillAllChildren(&menu, childrenMap)
				arr = append(arr, menu)
			}
		}
	}
	//存入缓存
	return &arr, err
}

// 获取管理员菜单数据
//
//	func (svc *MenuService) GenMenus(parent *model.SysMenu, menus []model.SysMenu) {
//		if parent.MenuType == "F" {
//			return
//		}
//		if parent.Children == nil {
//			//parent.Children = []model.SysMenu{}
//			parent.Children = make([]model.SysMenu, 0)
//		}
//		for i := range menus {
//			if menus[i].ParentId == parent.MenuId { //发现子菜单
//				parent.Children = append(parent.Children, menus[i])
//				svc.GenMenus(&menus[i], menus)
//			}
//		}
//		if len(parent.Children) == 0 {
//			return
//		}
//
// }
func (svc *MenuService) InitChildMap(menus []model.SysMenu) map[int64][]*model.SysMenu {
	childrenMap := make(map[int64][]*model.SysMenu)
	for i, _ := range menus {
		if menus[i].MenuType == "F" { //忽略按钮
			continue
		}
		//每个menu都预设子菜单项
		childrenMap[menus[i].MenuId] = make([]*model.SysMenu, 0)
	}

	for i, _ := range menus {
		if menus[i].MenuType == "F" { //忽略按钮
			continue
		}
		pid := menus[i].ParentId
		//组织父子关系
		childrenMap[pid] = append(childrenMap[pid], &menus[i])
		//svc.GenMenus(&menus[i], menus)
	}
	return childrenMap
}

// 检查菜单名是否唯一
func (svc *MenuService) CheckMenuNameUniqueAll(menuName string, parentId int64) bool {
	var dao dao.MenuDao
	return dao.CheckMenuNameExists(menuName, parentId)
}

// 检查菜单名是否唯一
func (svc *MenuService) CheckMenuNameUnique(menuName string, menuId, parentId int64) string {
	var dao dao.MenuDao
	exist := dao.CheckMenuNameExists(menuName, parentId)
	if exist {
		return "1"
	}
	return "0"
}

// 检查权限键是否唯一
func (svc *MenuService) CheckPermsUniqueAll(perms string) string {
	var dao dao.MenuDao
	entity, err := dao.CheckPermsUniqueAll(perms)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.MenuId > 0 {
		return "1"
	}
	return "0"
}

// 检查权限键是否唯一
func (svc *MenuService) CheckPermsUnique(perms string, menuId int64) string {
	var dao dao.MenuDao
	entity, err := dao.CheckPermsUniqueAll(perms)
	if err != nil {
		return "1"
	}
	if entity != nil && entity.MenuId > 0 && entity.MenuId != menuId {
		return "1"
	}
	return "0"
}

// 根据角色ID查询菜单
func (svc *MenuService) RoleMenuTreeData(roleId, userId int64) (*[]dto.Ztree, error) {
	var result *[]dto.Ztree
	menuList, err := svc.ListMenuNormalByUser(userId)
	if err != nil {
		return nil, err
	}
	var dao dao.MenuDao
	if roleId > 0 {
		roleMenuList, err := dao.SelectMenuTree(roleId)
		if err != nil || roleMenuList == nil {
			result, err = svc.InitZtree(menuList, nil, true)
		} else {
			result, err = svc.InitZtree(menuList, &roleMenuList, true)
		}
	} else {
		result, err = svc.InitZtree(menuList, nil, true)
	}

	return result, nil
}

// 对象转菜单树
func (svc *MenuService) InitZtree(menuList *[]model.SysMenu, roleMenuList *[]string, permsFlag bool) (*[]dto.Ztree, error) {
	var result []dto.Ztree
	isCheck := false
	if roleMenuList != nil && len(*roleMenuList) > 0 {
		isCheck = true
	}

	for _, obj := range *menuList {
		var ztree dto.Ztree
		ztree.Title = obj.MenuName
		ztree.Id = obj.MenuId
		ztree.Name = svc.transMenuName(obj.MenuName, permsFlag)
		ztree.Pid = obj.ParentId
		if isCheck {
			tmp := lv_conv.String(obj.MenuId) + obj.Perms
			tmpcheck := false
			for j := range *roleMenuList {
				if strings.Compare((*roleMenuList)[j], tmp) == 0 {
					tmpcheck = true
					break
				}
			}
			ztree.Checked = tmpcheck
		}
		result = append(result, ztree)
	}

	return &result, nil
}

func (svc *MenuService) transMenuName(menuName string, permsFlag bool) string {
	if permsFlag {
		return "<font color=\"#888\">&nbsp;&nbsp;&nbsp;" + menuName + "</font>"
	} else {
		return menuName
	}
}

func (svc *MenuService) IsRolePermited(roleKeys string, perms string) (bool, interface{}) {
	roles := strings.Split(roleKeys, ",")
	sql := "SELECT count(*) from sys_menu m,sys_role_menu rm,sys_role r where m.menu_id=rm.menu_id and rm.role_id = r.role_id and r.role_key in @Roles and m.perms=@Perms"
	count, err := lvdao.CountByNamedSql(sql, map[string]interface{}{"Roles": roles, "Perms": perms})
	return count > 0, err
}

func (svc *MenuService) FillAllChildren(m *model.SysMenu, childrenMap map[int64][]*model.SysMenu) {
	children := childrenMap[m.MenuId]
	if children == nil || len(children) <= 0 {
		return
	}
	for i := range children {
		m.Children = childrenMap[m.MenuId]
		svc.FillAllChildren(children[i], childrenMap)
	}
}
