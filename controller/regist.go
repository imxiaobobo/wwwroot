package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"project/dao"
)

//Regist 注册
func Regist(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		err := dao.SaveUser(username, password)
		if err != nil {
			fmt.Println("注册失败")
			t := template.Must(template.ParseFiles("view/Regist/regist.tpl"))
			t.Execute(w, "注册失败")
		} else {
			t := template.Must(template.ParseFiles("view/Login/login.tpl"))
			t.Execute(w, "注册成功")
		}
	} else {
		t := template.Must(template.ParseFiles("view/Regist/regist.tpl"))
		t.Execute(w, "请输入信息")
	}
}
