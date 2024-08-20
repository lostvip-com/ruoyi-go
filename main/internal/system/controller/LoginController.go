package controller

import (
	util2 "common/util"
	"github.com/gin-gonic/gin"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/utils/lv_conv"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_net"
	"github.com/lostvip-com/lv_framework/utils/lv_secret"
	"github.com/lostvip-com/lv_framework/web/lv_dto"
	"github.com/mojocn/base64Captcha"
	"github.com/mssola/user_agent"
	global2 "main/internal/common/global"
	"main/internal/common/myconf"
	"main/internal/system/model"
	"main/internal/system/service"

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
		verifyResult := base64Captcha.VerifyCaptcha(req.IdKey, req.ValidateCode)
		if !verifyResult {
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
		go func() {
			err = svc.SaveLoginLog2DB(token, user, ua, ip)
		}()
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
	logininfor.LoginLocation = lv_net.GetCityByIp(logininfor.Ipaddr)
	logininfor.Msg = msg
	logininfor.Status = "0"
	logininfor.Insert()
}

// 图形验证码
func (w *LoginController) CaptchaImage(c *gin.Context) {
	//config struct for digits
	//数字验证码配置
	//var configD = base64Captcha.ConfigDigit{
	//	Height:     80,
	//	Width:      240,
	//	MaxSkew:    0.7,
	//	DotCount:   80,
	//	CaptchaLen: 5,
	//}
	//config struct for audio
	//声音验证码配置
	//var configA = base64Captcha.ConfigAudio{
	//	CaptchaLen: 6,
	//	Language:   "zh",
	//}
	//config struct for Character
	//字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         4,
	}
	//创建声音验证码
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	//idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
	//以base64编码
	//base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
	//创建字符公式验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	//创建数字验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	//idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//以base64编码
	//base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)

	c.JSON(http.StatusOK, lv_dto.CaptchaRes{
		Code:           200,
		CaptchaEnabled: true,
		Uuid:           idKeyC,
		Img:            base64stringC,
		Msg:            "操作成功",
	})
}
