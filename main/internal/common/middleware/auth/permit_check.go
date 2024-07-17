package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/web/dto"
	"github.com/lostvip-com/lv_framework/web/router"
	"main/internal/common/myconf"
	"main/internal/system/service"
	"net/http"
	"strings"
)

// 鉴权中间件，只有登录成功之后才能通过
func PermitCheck(c *gin.Context) {
	//判断是否登录
	//根据url判断是否有权限
	url := c.Request.URL.Path
	strEnd := url[len(url)-1 : len(url)]
	if strings.EqualFold(strEnd, "/") {
		url = strings.TrimRight(url, "/")
	}
	//获取用户信息
	var userService service.UserService
	user := userService.GetProfile(c)
	if userService.IsAdmin(user.UserId) {
		c.Next()
		return
	}
	//获取权限标识
	permission := router.FindPermission(url)
	if permission == "" {
		c.Next()
		return
	}
	//获取用户菜单列表
	service := service.MenuService{}
	hasPermission, err := service.IsRolePermited(user.RoleKeys, permission)
	if err != nil {
		c.Redirect(http.StatusFound, myconf.GetConfigInstance().GetContextPath()+"/500")
		c.Abort()
		return
	}
	//fmt.Println("url:" + url + "---permission:" + permission)
	if !hasPermission {
		ajaxString := c.Request.Header.Get("X-Requested-With")
		if strings.EqualFold(ajaxString, "XMLHttpRequest") {
			c.JSON(http.StatusOK, dto.CommonRes{
				Code: 403,
				Msg:  "您没有操作权限",
			})
			c.Abort()
		} else {
			c.Redirect(http.StatusFound, myconf.GetConfigInstance().GetContextPath()+"/403")
			c.Abort()
		}
	}
	c.Next()
}
