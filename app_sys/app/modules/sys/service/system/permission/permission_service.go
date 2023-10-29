package permission

import (
	"html/template"
	"lostvip.com/utils/gconv"
	menuModel "robvi/app/modules/sys/model/system/menu"
	"robvi/app/modules/sys/service"
	menuService "robvi/app/modules/sys/service/system/menu"
	"strings"
)

// 根据用户id和权限字符串判断是否输出控制按钮
func GetPermiButton(u interface{}, permission, funcName, text, aclassName, iclassName string) template.HTML {

	result := HasPermi(u, permission)

	htmlstr := ""
	if result == "" {
		htmlstr = `<a class="` + aclassName + `" onclick="` + funcName + `" hasPermission="` + permission + `">
                    <i class="` + iclassName + `"></i> ` + text + `
                </a>`
	}

	return template.HTML(htmlstr)
}

// 根据用户id和权限字符串判断是否有此权限
func HasPermi(u interface{}, permission string) string {
	if u == nil {
		return "disabled"
	}

	uid := gconv.Int64(u)
	var userService service.UserService
	if uid <= 0 {
		return "disabled"
	}
	//获取权限信息
	var menus *[]menuModel.EntityExtend
	if userService.IsAdmin(uid) {
		menus, _ = menuService.SelectMenuNormalAll()
	} else {
		menus, _ = menuService.SelectMenusByUserId(gconv.String(uid))
	}

	if menus != nil && len(*menus) > 0 {
		for i := range *menus {
			if strings.EqualFold((*menus)[i].Perms, permission) {
				return ""
			}
		}
	}

	return "disabled"
}
