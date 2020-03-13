package main

import "shop/rabbit"

func main() {
	r := rabbit.NewRouting("test", "imooc_two")
	r.Consume()
}
