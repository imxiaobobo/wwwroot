package rabbit

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

/**
  simple模式生产者
*/
func (r *Rabbit) PublishSimple(msg string) {
	//申请队列,如果队列不存在则创建,保证队列存在,消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, //是否持久化
		false, //是否自动删除
		false, //是否具有排他性
		false, //是否阻塞
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false, //如果为true,会根据exchange类型和routkey规则,如果无法找到符合条件的队列,会把消息返还给发送者
		false, //如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者,会把消息返还给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
}

/**
  simple模式实例消费者
*/
func (r *Rabbit) ConsumeSimple() {
	//申请队列,如果队列不存在则创建,保证队列存在,消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, //是否持久化
		false, //是否自动删除
		false, //是否具有排他性
		false, //是否阻塞
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch := make(chan struct{})
	msg, err := r.channel.Consume(
		r.QueueName, //队列名
		"",          //区分多个消费者的
		true,        //是否自动应答
		false,       //是否具有排他性
		false,       //如果设置为true,表示不能将同一个conn中发送的消息传递给同个conn中的消费者
		false,       //队列是否阻塞
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for d := range msg {
			fmt.Println(string(d.Body))
		}
	}()
	log.Printf("等待接受消息...")
	<-ch
}
