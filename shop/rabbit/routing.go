/*
   @Time : 2020/3/13 09:54
   @Author : wangbo
   @File : routing
*/
package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
)

/**
  Routing模式:路由模式
  一个消息被多个消费者获取,并且消费的目标队列可被生产者指定
*/
type Routing struct {
	rabbit
}

//创建路由模式实例
func NewRouting(Exchange string, bangdingkey string) *Routing {
	r := newRabbitMq("", Exchange, bangdingkey) //传入的是交换机和bangdingkey
	s := &Routing{*r}
	return s
}

//发送者
func (r *Routing) Publish(msg string) {
	//1.创建交换机
	err := r.channel.ExchangeDeclare(r.exchange, "direct", true, false, false, false, nil)
	Err(err, "创建交换机失败")
	r.channel.Publish(r.exchange, r.bingdingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	})
}

//接受者
func (r *Routing) Consume() {
	//1.创建交换机
	err := r.channel.ExchangeDeclare(r.exchange, "direct", true, false, false, false, nil)
	Err(err, "创建交换机失败")
	//2.创建队列
	q, err := r.channel.QueueDeclare("", false, false, true, false, nil)
	Err(err, "创建队列失败")
	//绑定队列到exchange中
	err = r.channel.QueueBind(q.Name, r.bingdingKey, r.exchange, false, nil)
	Err(err, "绑定exchange失败")
	msg, err := r.channel.Consume(q.Name, "", true, false, false, false, nil)
	ch := make(chan struct{})
	go func() {
		for d := range msg {
			fmt.Println(string(d.Body))
		}
	}()
	<-ch
}
