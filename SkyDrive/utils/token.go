/*
   @Time : 2020/2/20 21:13
   @Author : wangbo
   @File : token
*/
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

const (
	token_salt = "_tokensalt"
)

func GetToken(username string) string {
	//username+timetemp+token_salt
	ts := fmt.Sprintf("%x", time.Now().Unix())
	md5 := md5.New()
	md5.Write([]byte(username + ts + token_salt))
	token := hex.EncodeToString(md5.Sum([]byte("")))
	return token + ts[:8]
}

//验证token的时效性
func IsTokenVaild(token string) bool {
	if len(token) != 40 {
		return false
	}
	ts := token[:8] //获取到token的后面八位十六进制时间
	t, err := strconv.ParseInt(ts, 16, 0)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if t < time.Now().Unix()-86400 {
		return false
	}
	return true
}
