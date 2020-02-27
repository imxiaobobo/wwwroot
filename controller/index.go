package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"project/dao"
	"strconv"
)

//Index 首页路由
func Index(w http.ResponseWriter, r *http.Request) {
	p := r.FormValue("pageNo")
	pageNo, err := strconv.ParseInt(p, 10, 0)
	if err != nil {
		fmt.Println(err)
		pageNo = 1
	}
	page, err := dao.PageBook(int(pageNo))
	if err != nil {
		fmt.Println(err)
		return
	}
	t := template.Must(template.ParseFiles("view/Index/index.tpl"))
	t.Execute(w, page)
}
