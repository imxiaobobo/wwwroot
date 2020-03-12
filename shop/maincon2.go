package main

import "shop/rabbit"

func main() {
	b := rabbit.Simple("test")
	b.Consume()
}
