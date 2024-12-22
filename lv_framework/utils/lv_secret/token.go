package lv_secret

import (
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/spf13/cast"
	"time"
)

type MyClaims struct {
	//【JWT ID】     该jwt的唯一ID编号
	//【issuer】     发布者的url地址
	//【issued at】  该jwt的发布时间；unix 时间戳
	//【subject】    该JWT所面向的用户，用于处理特定应用，不是常用的字段
	//【audience】   接受者的url地址
	//【expiration】 该jwt销毁的时间；unix时间戳
	//【not before】 该jwt的使用时间不能早于该时间；unix时间戳
	jwt.StandardClaims
	RefreshTime int64 //【The refresh time】 该jwt刷新的时间；unix时间戳
	//自定义
	UserId   int64
	TenantId int64
}

type Token struct {
	Claim    *MyClaims
	Token    string
	NewToken string
}

// timeout = 36000
// #刷新时间
// refresh = 18000
// #安全密钥robnote
// encryptKey = "cm9ibm90ZQ=="
// 创建Claims
func New(loginName string, userId, tenantId int64) *MyClaims {
	timeOut := cast.ToInt(lv_global.Config().GetVipperCfg().Get("token.timeout"))

	if timeOut <= 0 {
		timeOut = 3600
	}

	refresh := cast.ToInt(lv_global.Config().GetVipperCfg().Get("token.refresh"))

	if refresh <= 0 {
		refresh = timeOut / 2
	}

	var claims MyClaims
	claims.UserId = userId
	claims.TenantId = tenantId
	claims.Subject = loginName
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(timeOut)).Unix()
	claims.RefreshTime = time.Now().Add(time.Second * time.Duration(refresh)).Unix()
	return &claims
}

func (c *MyClaims) SetIss(issuer string) *MyClaims {
	c.StandardClaims.Issuer = issuer
	return c
}

func (c *MyClaims) SetSub(subject string) *MyClaims {
	c.StandardClaims.Subject = subject
	return c
}

func (c *MyClaims) SetAud(audience string) *MyClaims {
	c.StandardClaims.Audience = audience
	return c
}

func (c *MyClaims) SetNbf(notBefore int64) *MyClaims {
	c.StandardClaims.NotBefore = notBefore
	return c
}

func (c *MyClaims) Valid() error {
	//标准验证
	return c.StandardClaims.Valid()
}

// 创建token
func (claims *MyClaims) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encryptKey := cast.ToString(lv_global.Config().GetVipperCfg().Get("token.encryptKey"))
	mySignKeyBytes, err := base64.URLEncoding.DecodeString(encryptKey)
	if err != nil {
		return "", err
	}
	return token.SignedString(mySignKeyBytes)
}

// 验证token
func VerifyAuthToken(token string) (*Token, error) {
	encryptKey := cast.ToString(lv_global.Config().GetVipperCfg().Get("token.encryptKey"))
	mySignKeyBytes, err := base64.URLEncoding.DecodeString(encryptKey) //需要用和加密时同样的方式转化成对应的字节数组
	if err != nil {
		return nil, err
	}
	var myClaims MyClaims

	parseAuth, err := jwt.ParseWithClaims(token, &myClaims, func(*jwt.Token) (interface{}, error) {
		return mySignKeyBytes, nil
	})

	if err != nil {
		return nil, err
	}

	//验证claims
	if err := parseAuth.Claims.Valid(); err != nil {
		return nil, err
	}

	rs := new(Token)
	rs.Claim = &myClaims
	rs.Token = token
	//判断是否需要刷新
	if myClaims.RefreshTime > time.Now().Unix() {
		//生成新token
		newToken, err := New(myClaims.Subject, myClaims.UserId, myClaims.TenantId).CreateToken()
		if err == nil {
			rs.NewToken = newToken
		}
	}
	return rs, nil
}
