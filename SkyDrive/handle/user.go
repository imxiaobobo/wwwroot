/*
   @Time : 2020/2/14 19:57
   @Author : wangbo
   @File : user
*/
package handle

import (
	"SkyDrive/db"
	"SkyDrive/utils"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"html/template"
	"net/http"
)

//pwd_salt 盐
const (
	pwd_salt = "*#890"
)

//RegisterHandle 注册
func RegisterHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t := template.Must(template.ParseFiles("static/view/signup.tpl"))
		t.Execute(w, "")
	} else {
		// 获取用户名和密码,判断长短
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if len(username) < 3 || len(password) < 5 {
			w.Write([]byte("lnvalid parameter"))
			return
		}
		// 密码加密
		sha1 := sha1.New()
		sha1.Write([]byte(password + pwd_salt))
		passHash := hex.EncodeToString(sha1.Sum([]byte("")))
		// 通过用户名和密码完成入库
		suc := db.UserSignup(username, passHash)
		if suc {
			w.Write([]byte("success"))
		} else {
			w.Write([]byte("FAILED"))
		}
	}
}

//SignInHandle 登陆
func SignInHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("./static/view/login.tpl"))
		t.Execute(w, "")
	} else {
		//校验用户名密码
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		sha1 := sha1.New()
		sha1.Write([]byte(password + pwd_salt))
		encpassword := hex.EncodeToString(sha1.Sum([]byte("")))
		pwdCheck := db.UserSignIn(username, encpassword) //判断密码是否一致
		if !pwdCheck {
			data, _ := utils.NewResMsg(404, "密码不一致", nil).JsonBytes()
			w.Write(data)
			return
		}
		//生成token
		token := utils.GetToken(username)
		upRes := db.UpdateToken(username, token)
		if !upRes {
			data, _ := utils.NewResMsg(404, "生成token失败", nil).JsonBytes()
			w.Write(data)
			return
		}
		//登陆成功后重定向到首页
		resp := utils.NewResMsg(0, "OK", struct {
			Location, Username, Token string
		}{
			Location: "/user/index",
			Username: username,
			Token:    token,
		})
		rByte, err := resp.JsonBytes()
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write(rByte)
	}
}

////获取用户信息
//func UserInfoHandle(w http.ResponseWriter, r *http.Request) {
//
//	w.Write(respByte)
//}

//UserIndex 首页
func UserIndex(w http.ResponseWriter, r *http.Request) {
	//解析请求参数
	r.ParseForm()
	username := r.Form.Get("username")
	fmt.Println(username)
	//查询用户信息
	user, err := db.GetUserInfo(username)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	//组装并且响应用户
	resp := utils.NewResMsg(0, "OK", user)
	respByte, _ := resp.JsonBytes()
	fmt.Println(string(respByte))
	t := template.Must(template.ParseFiles("static/view/home.tpl"))
	t.Execute(w, string(respByte))
}
