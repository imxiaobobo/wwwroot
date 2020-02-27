/*
   @Time : 2020/1/13 15:12
   @Author : wangbo
   @File : db
*/
package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	dataBaseStr := "root:root@tcp(localhost:8889)/books"
	Db, err = sql.Open("mysql", dataBaseStr)
	if err != nil {
		panic(err)
	}
}
