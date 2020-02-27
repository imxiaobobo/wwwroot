/*
   @Time : 2020/1/13 14:37
   @Author : wang's
   @File : main
*/
package main

import (
	"net/http"
	"project/controller"
)

func main() {
	//静态文件
	{
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) //静态文件配置
	}

	//首页
	{
		http.HandleFunc("/", controller.Index) //首页
	}

	//user的登陆注册
	{
		http.HandleFunc("/login", controller.Login)                //登陆
		http.HandleFunc("/regist", controller.Regist)              //注册
		http.HandleFunc("/ajaxUserName", controller.CheckUserName) //ajax提交判断username
	}

	//book的操作
	{
		http.HandleFunc("/manager", controller.Manager)
		http.HandleFunc("/bookManager", controller.BookManager)
		http.HandleFunc("/addBook", controller.AddBook)
		http.HandleFunc("/bookDel", controller.BookDel)
		http.HandleFunc("/bookEdit", controller.BookEdit)
	}

	http.ListenAndServe(":8080", nil)

	/*r := route.NewMyMux()

	//通过调用一个方法来实现
	r.AddRoute("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Get!")
	})
	r.AddRoute("GET", "/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Get!")
	})

	r.AddRoute("GET", "/wo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Get!")
	})

	r.AddRoute("POST", "/hehe", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Get!")
	})

	r.Conn(":8080")*/

}
