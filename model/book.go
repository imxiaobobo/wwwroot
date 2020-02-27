/*
   @Time : 2020/1/16 09:39
   @Author : wangbo
   @File : book
*/
package model

type Book struct {
	ID      int     `json:"id"`
	Title   string  `json:"title"`
	Author  string  `json:"author"`
	Price   float64 `json:"price"`
	Sales   int     `json:"sales"`
	Stock   int     `json:"stock"`
	ImgPath string  `json:"img_path"`
}
