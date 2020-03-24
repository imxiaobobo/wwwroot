/*
   @Time : 2020/2/13 19:40
   @Author : wangbo
   @File : file
*/
package db

import (
	"SkyDrive/db/mysql"
	"database/sql"
	"fmt"
)

//TableFile 文件结构
type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

//OnFileUploadFinished 文件上传之后,将文件meta信息存入数据库
func OnFileUploadFinished(filehash, filename, fileaddr string, filesize int64) bool {
	stmt, err := mysql.DBConn().Prepare("insert ignore into tbl_file(file_sha1,file_name,file_size,file_addr,status) value (?,?,?,?,1)")
	if err != nil {
		fmt.Println("prepare:", err)
		return false
	}
	defer stmt.Close()
	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println("stmt:", err)
		return false
	}
	if rf, err := ret.RowsAffected(); err == nil {
		if rf <= 0 {
			fmt.Println("file with hash", err)
			return true
		}
		return false
	}
	return false
}

//GetFileMeta 根据filehash查询meta
func GetFileMeta(filehash string) (*TableFile, error) {
	stmt, err := mysql.DBConn().Prepare("select file_sha1,file_addr,file_name,file_size from tbl_file where file_sha1=? and status=1 limit 1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer stmt.Close()
	tfile := &TableFile{}
	err = stmt.QueryRow(filehash).Scan(&tfile.FileHash, &tfile.FileAddr, &tfile.FileName, &tfile.FileSize)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return tfile, nil
}
