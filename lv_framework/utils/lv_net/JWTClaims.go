// lv-auth - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2019, lostvip.org & hacpai.com
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package lv_net

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// BaseModel represents meta data of entity.
type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	Id          int64    `json:"id"`
	Phone       string   `json:"phone"`
	Username    string   `json:"username"`
	Realname    string   `json:"realname"`
	RoleCodes   []string `json:"permissions"`
	Permissions []string `json:"permissions"`
}

var (
	Secret     = "dong_tech" // 加盐
	ExpireTime = 3600        // token有效期
)

func GetJWTToken(username string, phone string) (signedToken string) {
	claims := &JWTClaims{
		Id:          1,
		Username:    username,
		Phone:       phone,
		RoleCodes:   []string{},
		Permissions: []string{},
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		panic(err)
	}
	return signedToken
}

func VerifyAction(strToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New("服务器繁忙")
	}

	claims, ok := token.Claims.(*JWTClaims)

	if !ok {
		fmt.Println("VerifyAction-----not ok--------------》claims:", claims, "---ok:", ok)
		return nil, errors.New("请重新登录")
	}

	fmt.Println("VerifyAction------ok-------------》claims:", claims, "---ok:", ok)

	if err := token.Claims.Valid(); err != nil {
		fmt.Println("token.Claims.Valid()-------------------》err:", err)
		return nil, errors.New("Token无效，请重新登录")
	}
	fmt.Println(" end VerifyAction-------------------》claims:", claims, "---error:", err)
	return claims, err
}
