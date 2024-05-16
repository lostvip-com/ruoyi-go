package service

import (
	"context"
	"errors"
	"github.com/mssola/user_agent"
	"lostvip.com/cache/myredis"
	"lostvip.com/utils/lv_net"
	"lostvip.com/utils/lv_secret"
	"main/app/system/model"
	"main/app/system/model/monitor/online"
	logininforService "main/app/system/service/monitor/logininfor"
	"strings"
	"time"
)

type SessionService struct{}

func (svc *SessionService) IsSignedIn(tokenStr string) bool {
	key := "login:" + tokenStr
	yes := myredis.GetInstance().Exists(context.Background(), key).Val()
	return yes > 0
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func (svc *SessionService) SignIn(loginnName, password string) (*model.SysUser, error) {
	//查询用户信息
	user := model.SysUser{LoginName: loginnName}
	err := user.FindOne()

	if err != nil {
		return nil, err
	}

	//校验密码
	pwdNew := user.LoginName + password + user.Salt

	pwdNew = lv_secret.MustEncryptString(pwdNew)

	if strings.Compare(user.Password, pwdNew) == -1 {
		return nil, errors.New("密码错误")
	}
	return &user, nil
}

// 清空用户菜单缓存

// 用户注销
func (svc *SessionService) SignOut(tokenStr string) error {
	myredis.GetInstance().Del(context.Background(), "login:"+tokenStr)
	entity := online.UserOnline{SessionId: "login:" + tokenStr}
	entity.Delete()
	return nil
}

// 强退用户
func (svc *SessionService) ForceLogout(token string) error {
	svc.SignOut(token)
	entity := online.UserOnline{SessionId: token}
	entity.Delete()
	return nil
}

func (svc *SessionService) SaveUserToSession(token string, user *model.SysUser, roleKeys string, userAgent string, ip string) error {
	// 保存用户信息到session
	loginLocation := lv_net.GetCityByIp(ip)
	//记录到redis
	redis := myredis.GetInstance()
	ctx := context.Background()
	fieldMap := make(map[string]interface{})
	fieldMap["userName"] = user.UserName
	fieldMap["userId"] = user.UserId
	fieldMap["loginName"] = user.LoginName
	fieldMap["avatar"] = user.Avatar
	fieldMap["ip"] = ip
	fieldMap["location"] = loginLocation
	fieldMap["roleKeys"] = roleKeys
	//其它
	//save to db
	ua := user_agent.New(userAgent)
	os := ua.OS()
	browser, _ := ua.Browser()
	//移除登录次数记录
	logininforService.RemovePasswordCounts(user.UserName)
	//
	var userOnline online.UserOnline
	userOnline.LoginName = user.UserName
	userOnline.Browser = browser
	userOnline.Os = os
	userOnline.DeptName = ""
	userOnline.Ipaddr = ip
	userOnline.ExpireTime = 1440
	userOnline.StartTimestamp = time.Now()
	userOnline.LastAccessTime = time.Now()
	userOnline.CreateTime = userOnline.StartTimestamp
	userOnline.Status = "on_line"
	userOnline.LoginLocation = loginLocation
	userOnline.SessionId = token
	err := userOnline.Delete()
	if err != nil {
		return err
	}
	err = userOnline.Save()
	if err != nil {
		return err
	}
	//
	key := "login:" + token
	err = redis.HMSet(ctx, key, fieldMap).Err()
	if err != nil {
		return err
	}
	err = redis.Expire(ctx, key, time.Hour).Err()
	if err != nil {
		return err
	}
	return err
}

func (svc *SessionService) Refresh(token string) {
	token = "login:" + token
	myredis.GetInstance().Expire(context.Background(), token, 8*time.Hour)
}
