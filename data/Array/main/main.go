/*
   @Time : 2020/2/28 20:59
   @Author : wangbo
   @File : main
*/
package main

import (
	"data/Array"
	"fmt"
)

func main() {
	arr := Array.NewArr(10)
	arr.AddLast(10)
	fmt.Println(arr)
}
