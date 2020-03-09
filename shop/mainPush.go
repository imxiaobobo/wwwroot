package main

import "shop/RabbitMq"

func main() {
	r := RabbitMq.NewRabbitMqSimple("test")
	r.PublishSimple("hello world")
}
