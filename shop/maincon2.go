package main

import "shop/rabbit"

func main() {
	r := rabbit.NewPub("test")
	r.Consume()
}
