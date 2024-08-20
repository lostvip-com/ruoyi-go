package controller

import (
	"common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/utils/lv_net"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"io"
	"main/internal/common/myconf"
	"main/internal/system/model"
	"main/internal/system/service"
	"net/http"
	"os"
)

type MainController struct{}

// 后台框架首页
func (w *MainController) Index(c *gin.Context) {
	w.goMain(c, "index")
	c.Abort()
}

func (w *MainController) goMain(c *gin.Context, indexPageDefault string) {
	var userService service.UserService
	user := userService.GetProfile(c)
	loginname := user.LoginName
	username := user.UserName
	avatar := user.Avatar
	if avatar == "" {
		avatar = "/resource/img/profile.jpg"
	}
	var menus *[]model.SysMenu
	//获取菜单数据
	menuService := service.MenuService{}
	if userService.IsAdmin(user.UserId) {
		tmp, err := menuService.SelectMenuNormalAll()
		if err == nil {
			menus = tmp
		}
	} else {
		tmp, err := menuService.SelectMenusByUserId(user.UserId)
		if err == nil {
			menus = tmp
		}
	}

	//获取配置数据
	var configService service.ConfigService
	sideTheme := configService.GetValueFromRam("sys.index.sideTheme")
	skinName := configService.GetValueFromRam("sys.index.skinName")
	//设置首页风格
	menuStyle := c.Query("menuStyle")
	cookie, _ := c.Request.Cookie("menuStyle")
	if cookie == nil {
		cookie = &http.Cookie{
			Name:     "menuStyle",
			Value:    menuStyle,
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
	}
	if menuStyle == "" { //未指定则从cookie中取
		menuStyle = cookie.Value
	}
	var targetIndex string         //默认首页
	if menuStyle == "index_left" { //指定了左侧风格,
		targetIndex = "index_left"
	} else { //否则默认风格
		targetIndex = indexPageDefault
	}
	//"menuStyle", cookie.Value, 1000, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly
	c.SetCookie(cookie.Name, menuStyle, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	util.BuildTpl(c, targetIndex).WriteTpl(gin.H{
		"avatar":    avatar,
		"loginname": loginname,
		"username":  username,
		"menus":     menus,
		"sideTheme": sideTheme,
		"skinName":  skinName,
	})
}

// 后台框架欢迎页面
func (w *MainController) Main(c *gin.Context) {
	util.BuildTpl(c, "main").WriteTpl()
}

// 下载 public/upload 文件头像之类
func (w *MainController) Download(c *gin.Context) {
	fileName := c.Query("fileName")
	//delete := c.Query("delete")
	if fileName == "" {
		util.BuildTpl(c, lv_dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	curDir, err := os.Getwd()
	filepath := curDir + "/static/upload/" + fileName
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		util.BuildTpl(c, lv_dto.ERROR_PAGE).WriteTpl(gin.H{
			"desc": "参数错误",
		})
		return
	}
	b, _ := io.ReadAll(file)
	c.Writer.Header().Add("Content-Disposition", "attachment")
	c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Writer.Write(b)
	c.Abort()
}

// 切换皮肤
func (w *MainController) SwitchSkin(c *gin.Context) {
	util.BuildTpl(c, "skin").WriteTpl()
}

// 注销
func (w *MainController) Logout(c *gin.Context) {
	var user service.SessionService
	tokenStr := lv_net.GetParam(c, "token")
	err := user.SignOut(tokenStr)
	if err != nil {
		return
	}
	path := myconf.GetConfigInstance().GetContextPath()
	c.SetCookie("token", "", -1, path, "", true, true)
	c.Redirect(http.StatusFound, "login")
	c.Abort()
}
