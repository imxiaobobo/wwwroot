/*
   @Time : 2020/2/23 20:01
   @Author : wangbo
   @File : userFile
*/
package db

import (
	"SkyDrive/db/mysql"
	"fmt"
	"time"
)

//UserFile 用户文件表
type UserFile struct {
	UserName, FileHash, FileName, UpdateAt, LastUpdate string
	FileSize                                           int64
}

//OnUserFileUploadFinished 更新用户文件表
func OnUserFileUploadFinished(username, filehash, filename string, filesize int64) bool {
	sql := "insert ignore into tbl_user_file(user_name,file_sha1,file_name,file_size,upload_at)values(?,?,?,?,?)"
	stmt, err := mysql.DBConn().Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return false
	}
	stmt.Close()
	_, err = stmt.Exec(username, filehash, filename, filesize, time.Now())
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
