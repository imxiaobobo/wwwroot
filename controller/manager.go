/*
   @Time : 2020/1/16 10:55
   @Author : wangbo
   @File : manager
*/
package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"project/dao"
	"project/model"
	"strconv"
)

func Manager(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("view/Manager/manager.tpl"))
	t.Execute(w, "")
}

//图书列表
func BookManager(w http.ResponseWriter, r *http.Request) {
	//获取页码
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
	t := template.Must(template.ParseFiles("view/Manager/book_manager.tpl"))
	t.Execute(w, page)
}

//图书删除
func BookDel(w http.ResponseWriter, r *http.Request) {
	bookId, _ := strconv.ParseInt(r.FormValue("bookid"), 10, 0)
	err := dao.DelBooks(int(bookId))
	if err != nil {
		http.Redirect(w, r, "/bookManager", http.StatusMovedPermanently)
	} else {
		http.Redirect(w, r, "/bookManager", http.StatusMovedPermanently)
	}
}

//增加
func AddBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// 获取图书信息
		title := r.PostFormValue("title")
		author := r.PostFormValue("author")
		price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
		sales, _ := strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
		stock, _ := strconv.ParseInt(r.PostFormValue("stock"), 10, 0)
		//创建book
		book := &model.Book{
			Title:   title,
			Author:  author,
			Price:   price,
			Sales:   int(sales),
			Stock:   int(stock),
			ImgPath: "staitc/img/default.jpg",
		}
		err := dao.AddBook(book)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/addBook", http.StatusMovedPermanently)
			return
		} else {
			http.Redirect(w, r, "/bookManager", http.StatusMovedPermanently)
		}
	} else {
		t := template.Must(template.ParseFiles("view/Manager/book_add.tpl"))
		t.Execute(w, "")
	}
}

//修改图书
func BookEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		bookid, _ := strconv.ParseInt(r.FormValue("bookid"), 10, 0)
		title := r.PostFormValue("title")
		author := r.PostFormValue("author")
		price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
		sales, _ := strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
		stock, _ := strconv.ParseInt(r.PostFormValue("stock"), 10, 0)
		//创建book
		book := &model.Book{
			ID:     int(bookid),
			Title:  title,
			Author: author,
			Price:  price,
			Sales:  int(sales),
			Stock:  int(stock),
		}
		err := dao.BookEdit(book)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/bookManager", http.StatusMovedPermanently)
		} else {
			http.Redirect(w, r, "/bookManager", http.StatusMovedPermanently)
		}
	} else {
		bookId, _ := strconv.ParseInt(r.FormValue("bookid"), 10, 0)
		b, _ := dao.GetBookById(int(bookId))
		t := template.Must(template.ParseFiles("view/Manager/book_edit.tpl"))
		t.Execute(w, b)
	}
}
