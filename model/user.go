/*
   @Time : 2020/1/13 15:12
   @Author : wangbo
   @File : user
*/
package model

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
