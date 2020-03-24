package main

import (
	"errors"
	"fmt"
	"reflect"
)

func Foo(a int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var s [10]int
	s[a] = 123
	fmt.Println(s)
}

func Test(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("除数不能为zero")
		return
	}
	res = a / b
	return
}

//闭包
func Seq() func() int {
	var a int
	return func() int {
		a++
		return a
	}
}

type People struct {
	Name string
}

func main() {
	var a interface{}
	a = 123

	reflectA := reflect.ValueOf(a)
	fmt.Println(reflectA.Int())

	//a := Seq()
	//val := a()
	//fmt.Println(val)
	//val = a()
	//fmt.Println(val)

	/*Foo(11)
	fmt.Println("程序继续执行了,因为recover已经捕获了异常,除非捕获代码中存在return")*/
	/*s, err := Test(1, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)*/
}
