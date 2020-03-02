package main

import (
	"SkyDrive/handle"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("服务器已启动")
	// 静态资源处理
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/file/upload", handle.UploadHandle) //上传
	//http.HandleFunc("/file/upload/suc", handle.UploadSucess)                     //上传成功
	http.HandleFunc("/file/meta", handle.GetFileMetaHandle)      //根据filehash获取meta
	http.HandleFunc("/file/query", handle.FileQueryHandle)       //根据数量获取meta切片
	http.HandleFunc("/file/download", handle.DownloadHandle)     //下载
	http.HandleFunc("/file/update", handle.FileUpdateMetaHandle) //修改文件名
	http.HandleFunc("/file/delete", handle.FileDelHandle)        //根据sha1删除

	http.HandleFunc("/user/register", handle.RegisterHandle) //注册
	http.HandleFunc("/user/login", handle.SignInHandle)      //登陆
	//http.HandleFunc("/user/index", handle.HTTPInterceptor(handle.UserIndex)) //主页
	//http.HandleFunc("/user/info", handle.HTTPInterceptor(handle.UserInfoHandle)) //用户信息

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器启动失败")
		return
	}
}
