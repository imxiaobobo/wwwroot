/*
   @Time : 2020/2/28 20:09
   @Author : wangbo
   @File : main
*/
package main

import (
	"data/Stack/ArrayStack"
	"fmt"
)

func main() {
	s := ArrayStack.NewStack(10)
	t := s.IsEmpty()
	i := s.GetSize()
	fmt.Println(t, i)
}
