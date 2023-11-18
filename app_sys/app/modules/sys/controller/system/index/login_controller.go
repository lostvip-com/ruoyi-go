package index

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/mssola/user_agent"
	"lostvip.com/myredis"
	"lostvip.com/utils/gconv"
	"lostvip.com/utils/ip"
	"lostvip.com/utils/lib_net"
	"lostvip.com/utils/lib_secret"
	"lostvip.com/utils/response"
	"net/http"
	"robvi/app/global"
	logininforModel "robvi/app/modules/sys/model/monitor/logininfor"
	"robvi/app/modules/sys/model/monitor/online"
	userModel "robvi/app/modules/sys/model/system/user"
	"robvi/app/modules/sys/service"
	logininforService "robvi/app/modules/sys/service/monitor/logininfor"
	"strings"
	"time"
)

type RegisterReq struct {
	UserName string `form:"username"  binding:"required,min=4,max=30"`
	Password string `form:"password" binding:"required,min=6,max=30"`
	//
	//ValidateCode string `form:"validateCode" binding:"min=4,max=10"`
	//IdKey        string `form:"idkey"        binding:"min=5,max=30"`

	ValidateCode string `form:"validateCode" `
	IdKey        string `form:"idkey" `
}

// 登陆页面
func Login(c *gin.Context) {

	if strings.EqualFold(c.Request.Header.Get("X-Requested-With"), "XMLHttpRequest") {
		response.ErrorResp(c).SetMsg("未登录或登录超时。请重新登录").WriteJsonExit()
		return
	}
	clientIp := lib_net.GetClientRealIP(c)
	errTimes := logininforService.GetPasswordCounts(clientIp)
	codeShow := 0 //默认不显示验证码
	if errTimes > 5 {
		codeShow = 1
	}
	response.BuildTpl(c, "login").WriteTpl(gin.H{
		"CodeShow": codeShow,
	})
}

// 验证登陆
func CheckLogin(c *gin.Context) {
	var req = RegisterReq{}
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	clientIp := lib_net.GetClientRealIP(c)
	errTimes4Ip := logininforService.GetPasswordCounts(clientIp)
	if errTimes4Ip > 5 { //超过5次错误开始校验验证码
		//比对验证码
		verifyResult := base64Captcha.VerifyCaptcha(req.IdKey, req.ValidateCode)
		if !verifyResult {
			response.ErrorResp(c).SetData(errTimes4Ip).SetMsg("验证码不正确").WriteJsonExit()
			return
		}
	}
	isLock := logininforService.CheckLock(req.UserName)
	if isLock {
		response.ErrorResp(c).SetMsg("账号已锁定，请30分钟后再试").WriteJsonExit()
		return
	}
	var userService service.UserService
	//验证账号密码
	user, err := userService.SignIn(req.UserName, req.Password)
	if err != nil {
		logininforService.SetPasswordCounts(clientIp)
		errTimes4UserName := logininforService.SetPasswordCounts(req.UserName)
		having := global.ErrTimes2Lock - errTimes4UserName
		SaveLogs(c, &req, "账号或密码不正确") //记录日志
		if having <= 5 {
			response.ErrorResp(c).SetData(errTimes4Ip).SetMsg("账号或密码不正确,还有" + gconv.String(having) + "次之后账号将锁定").WriteJsonExit()
		} else {
			response.ErrorResp(c).SetData(errTimes4Ip).SetMsg("账号或密码不正确!").WriteJsonExit()
		}
	} else {
		//保存在线状态
		cookie, _ := c.Request.Cookie("token")
		//token, _ := token.New(user.LoginName, user.UserId, user.TenantId).CreateToken()
		token := lib_secret.Md5(user.LoginName + time.UnixDate)
		maxage := 3600 * 8
		path := global.GetConfigInstance().GetContextPath()
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
		SaveUserToSession(token, user, c)
		SaveLogs(c, &req, "登陆成功") //记录日志
		response.SucessResp(c).SetData(token).SetMsg("登陆成功").WriteJsonExit()
	}
}

func SaveLogs(c *gin.Context, req *RegisterReq, msg string) {
	var logininfor logininforModel.Entity
	logininfor.LoginName = req.UserName
	logininfor.Ipaddr = c.ClientIP()
	userAgent := c.Request.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	logininfor.Os = ua.OS()
	logininfor.Browser, _ = ua.Browser()
	logininfor.LoginTime = time.Now()
	logininfor.LoginLocation = ip.GetCityByIp(logininfor.Ipaddr)
	logininfor.Msg = msg
	logininfor.Status = "0"
	logininfor.Insert()
}

// 保存用户信息到session
func SaveUserToSession(token string, user *userModel.SysUser, c *gin.Context) {
	loginIp := c.ClientIP()
	loginLocation := ip.GetCityByIp(loginIp)
	//记录到redis
	redis := myredis.GetInstance()
	ctx := context.Background()
	fieldMap := make(map[string]interface{})
	fieldMap["userName"] = user.UserName
	fieldMap["userId"] = user.UserId
	fieldMap["loginName"] = user.LoginName
	fieldMap["avatar"] = user.Avatar
	fieldMap["ip"] = loginIp
	fieldMap["location"] = loginLocation
	key := "login:" + token
	redis.HMSet(ctx, key, fieldMap)
	redis.Expire(ctx, key, time.Hour)
	//其它
	sessionId := user.UserId
	tmp, _ := json.Marshal(user)
	global.SessionList.Store(sessionId, string(tmp))
	//save to db
	userAgent := c.Request.Header.Get("User-Agent")
	ua := user_agent.New(userAgent)
	os := ua.OS()
	browser, _ := ua.Browser()

	//移除登陆次数记录
	logininforService.RemovePasswordCounts(user.UserName)
	//
	var userOnline online.UserOnline
	userOnline.LoginName = user.UserName
	userOnline.Browser = browser
	userOnline.Os = os
	userOnline.DeptName = ""
	userOnline.Ipaddr = loginIp
	userOnline.ExpireTime = 1440
	userOnline.StartTimestamp = time.Now()
	userOnline.LastAccessTime = time.Now()
	userOnline.Status = "on_line"
	userOnline.LoginLocation = loginLocation
	userOnline.Delete()
	userOnline.Insert()
}
