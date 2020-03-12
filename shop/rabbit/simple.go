package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Simple struct {
	rabbit
}

//获取simple模式实例
func NewSimple(QueueName string) *Simple {
	rabbit := newRabbitMq(QueueName, "", "")
	s := &Simple{*rabbit}
	return s
}

/**
  simple模式生产者
*/
func (s *Simple) Publish(msg string) {
	/**
	  申请队列,如果队列不存在则创建,保证队列存在,消息能发送到队列中
	  name 队列名
	  durable 是否持久化
	  autoDelete 是否自动删除
	  exclusive 是否具有排他性
	  noWait 是否阻塞
	*/
	_, err := s.channel.QueueDeclare(s.queueName, false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	/**
	  exchange 交换机
	  key 队列名
	  mandatory 如果为true,会根据exchange类型和routkey规则,如果无法找到符合条件的队列,会把消息返还给发送者
	  immediate 如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者,会把消息返还给发送者
	  msg 消息
	*/
	s.channel.Publish(s.exchange, s.queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	})
	fmt.Println("消息已发送")
}

/**
  simple模式实例消费者
*/
func (s *Simple) Consume() {
	//申请队列,如果队列不存在则创建,保证队列存在,消息能发送到队列中
	q, err := s.channel.QueueDeclare(s.queueName, false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch := make(chan struct{})
	/**
	  queue 队列名
	  consumer 区分多个消费者的
	  autoAck 是否自动应答
	  exclusive 是否具有排他性
	  noLocal 如果设置为true,表示不能将同一个conn中发送的消息传递给同个conn中的消费者
	  noWait 队列是否阻塞
	  args
	*/
	msg, err := s.channel.Consume(q.Name, "", true, false, false, false, nil)
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
