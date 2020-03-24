/*
   @Time : 2020/2/14 19:51
   @Author : wangbo
   @File : user
*/
package db

import (
	"SkyDrive/db/mysql"
	"database/sql"
	"fmt"
	"log"
)

//User 用户
type User struct {
	UserName, Email, Phone, SignupAt, LastActiveAt string
	Status                                         int
}

//UserSignup 通过用户名和密码完成入库
func UserSignup(username, password string) bool {
	stmt, err := mysql.DBConn().Prepare("insert ignore into tbl_user(user_name,user_pwd)value (?,?)")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()
	ret, err := stmt.Exec(username, password)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if re, err := ret.RowsAffected(); err == nil && re > 0 {
		return true
	}
	return false
}

//UserSignIn 判断密码是否一致
func UserSignIn(username, password string) bool {
	stmt, err := mysql.DBConn().Prepare("select * from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Println(err)
		return false
	} else if rows == nil {
		fmt.Println(username + "没这个人")
		return false
	}
	pRows := ParseRows(rows)
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == password {
		return true
	}
	return false
}

//ParseRows 查询字段返回map
func ParseRows(rows *sql.Rows) []map[string]interface{} {
	// 获取记录列(名)
	columns, _ := rows.Columns()
	// 创建列值的slice (values)，并为每一列初始化一个指针
	// scanArgs用作rows.Scan中的传入参数
	scanArgs := make([]interface{}, len(columns))
	for index := range scanArgs {
		var temp interface{}
		scanArgs[index] = &temp
	}
	// record为每次迭代中存储行记录的临时变量
	record := make(map[string]interface{})
	// records为函数最终返回的数据(列表)
	records := make([]map[string]interface{}, 0)
	// 迭代行记录
	for rows.Next() {
		err := rows.Scan(scanArgs...) //每Scan一次，将一行数据保存到record字典
		checkErr(err)
		for i, col := range scanArgs {
			if col != nil {
				record[columns[i]] = *col.(*interface{})
			}
		}
		records = append(records, record)
	}
	return records
}

//checkErr 错误处理
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

//UpdateToken token入库
func UpdateToken(username, token string) bool {
	stmt, err := mysql.DBConn().Prepare("replace into tbl_user_token(user_name,user_token) values(?,?)")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//GetUserInfo 获取用户信息
func GetUserInfo(username string) (u *User, err error) {
	u = &User{}
	stmt, err := mysql.DBConn().Prepare("select user_name,signup_at from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(username).Scan(&u.UserName, &u.SignupAt)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
