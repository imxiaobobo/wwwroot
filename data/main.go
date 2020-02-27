/*
   @Time : 2020/2/23 13:55
   @Author : wangbo
   @File : main
*/
package main

import (
	"data/Array"
	"fmt"
)

func main() {
	a := Array.NewArr(3)
	a.AddLast(5)
	a.AddLast(4)
	a.AddLast("wang")
	a.AddLast(2)
	fmt.Println(a)
}
