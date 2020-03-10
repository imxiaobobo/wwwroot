package main

import (
	"fmt"
)

func main() {
	fmt.Println(main) //打印main的话,打印的是main函数所在内存的内存地址
	//基于进制定义整形
	var a = 123   //十进制
	var b = 0123  //八进制
	var c = 0x123 //十六进制
	var d float64
	d = float64(8) / float64(3)

	fmt.Printf("%d,%o,%x", a, b, c)
	fmt.Println()
	fmt.Printf("%.2f", d)
	fmt.Println()
	var s string
	//s := '中'
	//x := "中"
	//var s []byte  uint8
	//s := []rune{20013, 20013, 20013, 20013}
	fmt.Printf("%v", s)
	//string []byte(20013)         rune int32
}
