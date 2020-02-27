package handle

import (
	"SkyDrive/db"
	"SkyDrive/meta"
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//UploadHandle 上传
func UploadHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("static/view/index.tpl"))
		t.Execute(w, "")
	} else {
		//上传逻辑
		file, head, err := r.FormFile("file") //获取file
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fileMeta := meta.FileMeta{ //组装filemeta
			FileSha1: "",
			FileName: head.Filename,
			Location: "./temp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		f, err := os.OpenFile(fileMeta.Location, os.O_CREATE|os.O_RDWR, 0755) //打开文件
		bf := bufio.NewWriter(f)                                              //写缓存
		if err != nil {
			fmt.Println("f", err)
			return
		}
		defer f.Close()
		fileMeta.FileSize, err = io.Copy(bf, file) //copy写缓存,复制文件,并获得filesize
		if err != nil {
			fmt.Println("copy", err)
			return
		}
		filehash := meta.FileHash(f) //sha1算法
		fileMeta.FileSha1 = filehash //存入fileMeta

		r.ParseForm()
		username := r.Form.Get("username")
		suc := db.OnUserFileUploadFinished(username, fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize)
		if !suc {
			w.Write([]byte("Upload Failed"))
			return
		}
		_ = meta.UpdateFileMetaDB(fileMeta) //将文件原信息存入数据库
		http.Redirect(w, r, "/user/index", http.StatusFound)
	}
}

////UploadSucess 上传成功
//func UploadSucess(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "uploadFile Sucess")
//}

//GetFileMetaHandle 获取Meta
func GetFileMetaHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileHash := r.Form["filehash"][0]
	//fMeta := meta.GetFileMeta(fileHash)
	fMeta, err := meta.GetFileMetaDB(fileHash)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := json.Marshal(fMeta)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

//FileQueryHandle 获取多个meta
func FileQueryHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	limitCnt, _ := strconv.Atoi(r.Form.Get("limit"))
	fileMeatas := meta.GetLastFileMetas(limitCnt)
	data, err := json.Marshal(fileMeatas)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(data)
}

//DownloadHandle 下载功能
func DownloadHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	//fm := meta.GetFileMeta(fsha1)
	fm, err := meta.GetFileMetaDB(fsha1)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Open(fm.Location)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-Disposition", "attachment;filename="+fm.FileName)
	_, err = w.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//FileUpdateMetaHandle 修改原信息接口
func FileUpdateMetaHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	opType := r.Form.Get("op")
	filehash := r.Form.Get("filehash")
	filename := r.Form.Get("filename")
	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//先获取元信息
	//curmeta := meta.GetFileMeta(filehash)
	curmeta, err := meta.GetFileMetaDB(filehash)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(curmeta)
	curmeta.FileName = filename
	s1 := strings.Split(curmeta.Location, ".")
	s := "./temp/" + filename + "." + s1[2]
	err = os.Rename(curmeta.Location, s)
	if err != nil {
		fmt.Println(err)
		return
	}
	meta.UploadFileMeta(*curmeta)
	data, err := json.Marshal(curmeta)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

//FileDelHandle 删除
func FileDelHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form.Get("filehash")
	fMeta, err := meta.GetFileMetaDB(filehash)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Remove(fMeta.Location)
	err = meta.RemoveFileMeta(filehash)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}
