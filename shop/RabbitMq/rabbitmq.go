package RabbitMq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

/**
  rabbitmq simple 模式
*/
const MQURL = "amqp://guest:guest@127.0.0.1:5672/"

type RabbitMq struct {
	conn                            *amqp.Connection
	channel                         *amqp.Channel
	QueueName, Exchange, Key, Mqurl string
}

/**
  创建结构体实例
*/
func newRabbitMq(QueueName, Exchange, Key string) *RabbitMq {
	rabbit := &RabbitMq{
		QueueName: QueueName,
		Exchange:  Exchange,
		Key:       Key,
		Mqurl:     MQURL,
	}
	//创建连接
	var err error
	rabbit.conn, err = amqp.Dial(rabbit.Mqurl) //创建一个连接
	rabbit.failOnErr(err, "dial错误")
	rabbit.channel, err = rabbit.conn.Channel() //创建一个channel
	rabbit.failOnErr(err, "获取channel失败")
	return rabbit
}

/**
  断开连接
*/
func (r *RabbitMq) Close() {
	r.channel.Close()
	r.conn.Close()
}

/**
  错误处理
*/
func (r *RabbitMq) failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

/**
  创建简单工作模式
*/
func NewRabbitMqSimple(queueName string) *RabbitMq {
	rabbit := newRabbitMq(queueName, "", "")
	return rabbit
}

/**
  简单模式下生产者
*/
func (r *RabbitMq) PublishSimple(message string) {
	//申请队列,如果队列不存在,会自动创建,存在则跳过创建
	_, err := r.channel.QueueDeclare(r.QueueName, false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送消息到队列中
	r.channel.Publish(r.Exchange, r.QueueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})
	fmt.Println("发送成功")
}

/**
  简单模式消费消息
*/
func (r *RabbitMq) ConsumeSimple() {
	//申请队列
	_, err := r.channel.QueueDeclare(r.QueueName, false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//接受消息
	message, err := r.channel.Consume(r.QueueName, "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
	}
	temp := make(chan struct{})
	go func() {
		for d := range message {
			fmt.Println(string(d.Body))
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-temp
}
