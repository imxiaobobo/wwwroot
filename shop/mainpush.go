package main

import (
	"fmt"
	"shop/rabbit"
	"strconv"
	"time"
)

func main() {
	r := rabbit.NewPub("test")
	for i := 0; i <= 100; i++ {
		r.Publish("hello world" + strconv.Itoa(i))
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}
