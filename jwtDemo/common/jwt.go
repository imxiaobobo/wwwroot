package common

import (
	"github.com/dgrijalva/jwt-go"
	"jwtDemo/model"
	"log"
	"time"
)

var jwtKey = []byte("a_secret_crect") //创建jwt秘钥

type Claims struct {
	UserId             uint
	jwt.StandardClaims //继承
}

func ReleaseToken(user model.UserModel) (tokenStr string, err error) {
	clims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(), //设置token过期时间7天
			IssuedAt:  time.Now().Unix(),                         //token发放的时间
			Issuer:    "oceanlearn.tech",                         //token发放人
			Subject:   "user token",                              //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clims)
	tokenStr, err = token.SignedString(jwtKey)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
