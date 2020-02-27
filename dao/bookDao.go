/*
   @Time : 2020/1/16 09:53
   @Author : wangbo
   @File : bookDao
*/
package dao

import (
	"fmt"
	"log"
	"math"
	"project/model"
	"project/util"
)

/**
  获取所有图书
*/
func GetAllBooks() ([]*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	rows, err := util.Db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		var book = &model.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

/**
  删除图书
*/
func BookDel(bookId int) (err error) {
	sqlStr := "delete from books where id=?"
	_, err = util.Db.Exec(sqlStr, bookId)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

/**
  添加图书
*/
func AddBook(b *model.Book) (err error) {
	sqlStr := "insert into books(title,author,price,sales,stock,img_path)values(?,?,?,?,?,?)"
	_, err = util.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

/**
  删除图书
*/
func DelBooks(bookid int) (err error) {
	sqlStr := "delete from books where id=?"
	_, err = util.Db.Exec(sqlStr, bookid)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

/**
  修改图书
*/
func BookEdit(b *model.Book) (err error) {
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	_, err = util.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

/**
  根据id获取图书信息
*/
func GetBookById(bookid int) (*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id=?"
	b := &model.Book{}
	err := util.Db.QueryRow(sqlStr, bookid).Scan(&b.ID, &b.Title, &b.Author, &b.Price, &b.Sales, &b.Stock, &b.ImgPath)
	if err != nil {
		return nil, err
	}
	return b, nil
}

/**
  分页函数
*/
func PageBook(pageNo int) (*model.Page, error) {
	//获取数据库中图书的总数
	sqlStr := "select count(*) from books"
	var total int64
	page := &model.Page{}
	util.Db.QueryRow(sqlStr).Scan(&total)
	//设置每页显示几条
	pageSize := 4
	count := int64(math.Ceil(float64(total) / float64(pageSize)))
	//获取当前页的图书
	sqlStr = "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	rows, err := util.Db.Query(sqlStr, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var book = &model.Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			return nil, err
		}
		page.Book = append(page.Book, book)
	}
	page.PageNo = int64(pageNo)
	page.PageSize = int64(pageSize)
	page.Count = count
	page.Total = total
	return page, nil
}
