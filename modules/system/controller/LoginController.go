package controller

import (
	global2 "common/global"
	"common/myconf"
	util2 "common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_net"
	"github.com/lostvip-com/lv_framework/utils/lv_secret"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"github.com/mssola/user_agent"
	"net/url"
	"system/model"
	"system/service"

	"net/http"
	"strings"
	"time"
)

type LoginController struct {
}
type RegisterReq struct {
	UserName string `form:"username"  binding:"required,min=4,max=30"`
	Password string `form:"password" binding:"required,min=6,max=30"`
	//
	//ValidateCode string `form:"validateCode" binding:"min=4,max=10"`
	//IdKey        string `form:"idkey"        binding:"min=5,max=30"`

	ValidateCode string `form:"validateCode" `
	IdKey        string `form:"idkey" `
}

// 登录页面
func (w *LoginController) Login(c *gin.Context) {

	if strings.EqualFold(c.Request.Header.Get("X-Requested-With"), "XMLHttpRequest") {
		util2.ErrorResp(c).SetMsg("未登录或登录超时。请重新登录").WriteJsonExit()
		return
	}
	clientIp := lv_net.GetClientRealIP(c)
	var logininforService service.LoginInforService
	errTimes := logininforService.GetPasswordCounts(clientIp)
	codeShow := 0 //默认不显示验证码
	if errTimes > 5 {
		codeShow = 1
	}
	util2.BuildTpl(c, "login").WriteTpl(gin.H{
		"CodeShow": codeShow,
	})
}

/**
 * 调用公共服务提供的验证方法
 */
func (w *LoginController) VerifyCaptcha(IdKey string, ValidateCode string) (string, error) {
	return "", nil
}

// 验证登录
func (w *LoginController) CheckLogin(c *gin.Context) {
	var req = RegisterReq{}
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		util2.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	clientIp := lv_net.GetClientRealIP(c)
	var logininforService service.LoginInforService
	errTimes4Ip := logininforService.GetPasswordCounts(clientIp)
	if errTimes4Ip > 5 { //超过5次错误开始校验验证码
		//比对验证码
		_, err := w.VerifyCaptcha(req.IdKey, req.ValidateCode)
		if err != nil {
			util2.ErrorResp(c).SetData(errTimes4Ip).SetMsg("验证码不正确").WriteJsonExit()
			return
		}
	}
	isLock := logininforService.CheckLock(req.UserName)
	if isLock {
		util2.ErrorResp(c).SetMsg("账号已锁定，请30分钟后再试").WriteJsonExit()
		return
	}
	var userService service.UserService
	//验证账号密码
	user, err := userService.SignIn(req.UserName, req.Password)
	if err != nil {
		logininforService.SetPasswordCounts(clientIp)
		errTimes4UserName := logininforService.SetPasswordCounts(req.UserName)
		having := global2.ErrTimes2Lock - errTimes4UserName
		w.SaveLogs(c, &req, "账号或密码不正确") //记录日志
		if having <= 5 {
			util2.ErrorResp(c).SetData(errTimes4Ip).SetMsg("账号或密码不正确,还有" + lv_conv.String(having) + "次之后账号将锁定").WriteJsonExit()
		} else {
			util2.ErrorResp(c).SetData(errTimes4Ip).SetMsg("账号或密码不正确!").WriteJsonExit()
		}
	} else {
		//保存在线状态
		cookie, _ := c.Request.Cookie("token")
		//token, _ := token.New(user.LoginName, user.UserId, user.TenantId).CreateToken()
		token := lv_secret.Md5(user.LoginName + time.UnixDate)
		maxage := 3600 * 8
		path := myconf.GetConfigInstance().GetContextPath()
		if cookie == nil {
			cookie = &http.Cookie{
				Path:     path,
				Name:     "token",
				MaxAge:   maxage,
				Value:    token,
				HttpOnly: true,
			}
			http.SetCookie(c.Writer, cookie)
		}
		c.SetCookie(cookie.Name, token, maxage, path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		// 生成token
		ua := c.Request.Header.Get("User-Agent")
		roles, err := userService.GetRoleKeys(user.UserId)
		lv_err.HasErrAndPanic(err)
		var svc service.SessionService
		ip := lv_net.GetClientRealIP(c)
		err = svc.SaveUserToSession(token, user, roles)
		//go func() {
		err = svc.SaveLoginLog2DB(token, user, ua, ip)
		//}()
		if err != nil {
			lv_log.Error(err)
			lv_err.PrintStackTrace(err)
			w.SaveLogs(c, &req, "登录失败！"+err.Error()) //记录日志
			util2.Fail(c, "登录失败")
		} else {
			w.SaveLogs(c, &req, "login success") //记录日志
			util2.Success(c, token, "login success!")
		}
	}
}

func (w *LoginController) SaveLogs(c *gin.Context, req *RegisterReq, msg string) {
	var logininfor model.SysLoginInfo
	logininfor.LoginName = req.UserName
	logininfor.Ipaddr = c.ClientIP()
	userAgent := c.Request.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	logininfor.Os = ua.OS()
	logininfor.Browser, _ = ua.Browser()
	logininfor.LoginTime = time.Now()
	logininfor.LoginLocation = util2.GetCityByIp(logininfor.Ipaddr)
	logininfor.Msg = msg
	logininfor.Status = "0"
	logininfor.Insert()
}

// CaptchaImage 图形验证码生成逻辑,使用其它服务的公共接口生成，不再单独维护
func (w *LoginController) CaptchaImage(c *gin.Context) {
	//传参数
	clientId := c.PostForm("uuid")
	//返回值
	params := url.Values{}
	params.Set("uuid", clientId)
	url_captcha := lv_global.Config().GetValueStr("url_captcha")
	if url_captcha == "" {
		util2.Err(c, "请配置验证码地址url_captcha")
		return
	}
	json, err := lv_net.PostForm(url_captcha, params)
	lv_err.HasErrAndPanic(err)
	dataMap := lv_conv.ToMap(json)
	base64stringC := dataMap["data"]
	c.JSON(http.StatusOK, lv_dto.CaptchaRes{
		Code:           200,
		CaptchaEnabled: true,
		Uuid:           clientId,
		Img:            base64stringC,
		Msg:            "操作成功",
	})
}
