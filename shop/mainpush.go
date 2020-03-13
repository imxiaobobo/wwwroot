package main

import (
	"shop/rabbit"
	"strconv"
)

func main() {
	r1 := rabbit.NewRouting("test", "imooc_one")
	r2 := rabbit.NewRouting("test", "imooc_two")

	for i := 0; i <= 100; i++ {
		r1.Publish("one" + strconv.Itoa(i))
		r2.Publish("two" + strconv.Itoa(i))
	}
}
