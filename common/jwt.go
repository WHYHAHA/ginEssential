package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"why/ginessential/model"
)

var jwtkey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User)	(string, error)  {
	expirationTime := time.Now().Add(7*24*time.Hour)
	claims := &Claims{
		UserId:         user.ID,
		StandardClaims: jwt.StandardClaims{
			// 到期时间
			ExpiresAt:expirationTime.Unix(),
			//发放时间
			IssuedAt: time.Now().Unix(),
			//发布者
			Issuer: "oceanLearn.tech",
			//主题
			Subject: "user token",
		},
	}
	//写入秘钥
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err := token.SignedString(jwtkey)

	if err != nil {
		return "",err
	}

	return tokenString,nil
}
// token 由三部分组成 第一部分 协议头header(使用的加密协议) 第二部分 荷载payload(储存的是claims对象中的所有信息) 第三部分 签证signature(是由jwtkey加前面两部分哈希的值)
//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIsImV4cCI6MTYyNDAwNzQ4OCwiaWF0IjoxNjIzNDAyNjg4LCJpc3MiOiJvY2VhbkxlYXJuLnRlY2giLCJzdWIiOiJ1c2VyIHRva2VuIn0.tBgpxvXhnrNzfEGo0UtXxt7hL12f0cestiYLfLHVuBk

func ParseToken(tokenString string) (*jwt.Token,*Claims,error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})

	return token,claims,err
}