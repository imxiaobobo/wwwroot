/*
   @Time : 2020/1/14 10:54
   @Author : wangbo
   @File : ajax
*/
package controller

import (
	"database/sql"
	"net/http"
	"project/util"
)

//CheckUserName 判断username
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	sqlStr := "select username from user where username=?"
	err := util.Db.QueryRow(sqlStr, username).Scan()
	if err != nil {
		if err != sql.ErrNoRows {
			w.Write([]byte("用户名已存在"))
			return
		}
	}
}
