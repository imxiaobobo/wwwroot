package common

import (
	"github.com/dgrijalva/jwt-go"
	"jwtDemo/model"
	"log"
	"time"
)

var jwtKey = []byte("a_secret_create") //创建jwt秘钥

// Claims jwt结构体
type Claims struct {
	UserID             uint
	jwt.StandardClaims //继承
}

/**
  生成jwt
*/
func ReleaseToken(user model.UserModel) (tokenStr string, err error) {
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(), //设置token过期时间7天
			IssuedAt:  time.Now().Unix(),                         //token发放的时间
			Issuer:    "oceanlearn.tech",                         //token发放人
			Subject:   "user token",                              //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(jwtKey)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

/**
  分解jwt
*/
func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
