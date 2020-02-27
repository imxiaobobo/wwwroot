package controller

import (
	"html/template"
	"net/http"
	"project/dao"
)

//Login 登陆
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		user, err := dao.Login(username, password)
		if err != nil {
			t := template.Must(template.ParseFiles("view/Login/login.tpl"))
			t.Execute(w, err.Error())
			return
		}
		t := template.Must(template.ParseFiles("view/Login/login_success.tpl"))
		t.Execute(w, user)
	} else {
		t := template.Must(template.ParseFiles("view/Login/login.tpl"))
		t.Execute(w, "请输入用户名和密码")
	}
}
