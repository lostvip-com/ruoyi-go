package service

import (
	"github.com/spf13/cast"
	"html/template"
	dao2 "main/internal/system/dao"
	"main/internal/system/model"
	"strings"
)

type PermissionService struct{}

// 根据用户id和权限字符串判断是否输出控制按钮
func (svc *PermissionService) GetPermiButton(u interface{}, permission, funcName, text, aclassName, iclassName string) template.HTML {
	result := svc.HasPermi(u, permission)

	htmlstr := ""
	if result == "" {
		htmlstr = `<a class="` + aclassName + `" onclick="` + funcName + `" hasPermission="` + permission + `">
                    <i class="` + iclassName + `"></i> ` + text + `
                </a>`
	}

	return template.HTML(htmlstr)
}

// 根据用户id和权限字符串判断是否有此权限
func (svc *PermissionService) HasPermi(u interface{}, permission string) string {
	if u == nil {
		return "disabled"
	}

	uid := cast.ToInt64(u)
	var userService UserService
	if uid <= 0 {
		return "disabled"
	}
	//获取权限信息
	var menus []model.SysMenu
	var dao dao2.MenuDao
	if userService.IsAdmin(uid) {
		menus, _ = dao.SelectMenuNormalAll()
	} else {
		menus, _ = dao.SelectMenusByUserId(uid)
	}

	if menus != nil && len(menus) > 0 {
		for i := range menus {
			if strings.EqualFold((menus)[i].Perms, permission) {
				return ""
			}
		}
	}

	return "disabled"
}
