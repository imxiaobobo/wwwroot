/*
   @Time : 2020/1/17 15:08
   @Author : wangbo
   @File : main
*/
package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}
func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil) //传入实现了ServeHTTP方法的对象
}
