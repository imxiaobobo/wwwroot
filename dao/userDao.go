/*
   @Time : 2020/1/13 15:21
   @Author : wangbo
   @File : userDao
*/
package dao

import (
	"database/sql"
	"fmt"
	"project/model"
	"project/util"
)

/**
  登陆方法
*/
func Login(username, password string) (*model.User, error) {
	sqlStr := "select id,username,password from user where username=? and password=?"
	stmtSel, err := util.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer stmtSel.Close()
	user := &model.User{}
	err = stmtSel.QueryRow(username, password).Scan(&user.ID, &user.UserName, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
	}
	return user, nil
}

/**
  注册
*/
func SaveUser(username, password string) (err error) {
	sqlStr := "select id from user where username=?"
	stmtSel, err := util.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = stmtSel.QueryRow(username).Scan()
	if err != sql.ErrNoRows {
		fmt.Println("用户已存在")
		return
	}
	sqlStr = "insert into user (username,password) values(?,?)"
	stmtIns, err := util.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = stmtIns.Exec(username, password)
	if err != nil {
		fmt.Println("注册失败")
		return
	}
	return
}
