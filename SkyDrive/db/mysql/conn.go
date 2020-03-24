/*
   @Time : 2020/2/11 20:20
   @Author : wangbo
   @File : conn
*/
package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/skydrive?charset=utf8")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("mysql conn err+", err.Error())
		os.Exit(1)
	}
}

func DBConn() *sql.DB {
	return db
}
