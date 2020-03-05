package meta

import (
	"SkyDrive/db"
	"errors"
	"fmt"
	"sort"
)

//FileMeta 文件元信息结构
type FileMeta struct {
	FileSha1 string //文件哈希
	FileName string //文件名
	FileSize int64  //文件大小
	Location string //文件路径
	UploadAt string //上传时间
}

// 文件原信息的全局变量
var fileMetas map[string]FileMeta

//init 初始化全局文件原信息
func init() {
	fileMetas = make(map[string]FileMeta)
}

//UploadFileMeta 将文件的sha1信息作为key,原信息结构体作为value存入全局文件原信息
func UploadFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

func UpdateFileMetaDB(fmeta FileMeta) bool {
	return db.OnFileUploadFinished(fmeta.FileSha1, fmeta.FileName, fmeta.Location, fmeta.FileSize)
}

//GetFileMeta 根据sha1获取文件原信息
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

func GetFileMetaDB(filehash string) (*FileMeta, error) {
	tfile, err := db.GetFileMeta(filehash)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmeta := &FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}
	return fmeta, nil
}

//GetLastFileMetas 根据获取个数获取文件源信息
func GetLastFileMetas(count int) []FileMeta {
	var fMetaArray []FileMeta
	for _, v := range fileMetas {
		fMetaArray = append(fMetaArray, v)
		fmt.Println(fMetaArray)
	}
	sort.Sort(ByUploadTime(fMetaArray))
	if count > len(fileMetas) {
		count = len(fileMetas)
	}
	return fMetaArray[:count]
}

func RemoveFileMeta(sha1 string) (err error) {
	_, ok := fileMetas[sha1]
	if !ok {
		return errors.New("无此fileHash")
	}
	delete(fileMetas, sha1)
	return
}
