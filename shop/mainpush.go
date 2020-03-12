package main

import (
	"fmt"
	"shop/rabbit"
	"strconv"
	"time"
)

func main() {
	b := rabbit.New("test", 1)
	for i := 0; i < 100; i++ {
		b.PublishSimple("hello world" + strconv.Itoa(i))
		time.Sleep(time.Second * 1)
		fmt.Println(i)
	}
}
