package rabbit

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

/**
  virtual host:有效的对数据进行逻辑上的隔离
  connections:连接
  Exchange:交换机,当生产者发送消息的时候,会传入到交换机中,根据规则会将数据绑定到不同的key中
  channel:通道
  queue:队列
  bingding:将队列绑定到不同的交换机上,实现不同的工作模式
*/

/**
  rabbit的连接信息
  格式:amqp://账号:密码@地址:端口号/virtualHost
*/
const URL = "amqp://guest:guest@127.0.0.1:5672/imooc" //rabbit的连接信息

type Rabbit struct {
	conn        *amqp.Connection //rabbit连接
	channel     *amqp.Channel    //rabbit管道
	QueueName   string           //队列
	Exchange    string           //交换机
	BingdingKey string
	Url         string
}

//创建rabbit基础实例
func newRabbitMq(QueueName, Exchange, BingdingKey string) *Rabbit {
	r := &Rabbit{
		QueueName:   QueueName,
		Exchange:    Exchange,
		BingdingKey: BingdingKey,
		Url:         URL,
	}
	var err error
	r.conn, err = amqp.Dial(r.Url) //传入连接信息
	r.Err(err, "连接失败")
	r.channel, err = r.conn.Channel() //获取管道
	r.Err(err, "获取管道失败")
	return r
}

//关闭连接
func (r *Rabbit) Close() {
	r.channel.Close() //先关闭管道
	r.conn.Close()    //在关闭连接
}

//错误处理
func (r *Rabbit) Err(err error, msg string) {
	if err != nil {
		log.Fatalln(err, msg)
		panic(fmt.Sprintln(err, msg))
	}
}

/**
  模式实例
  types 1. simple 2. 订阅者模式
*/
func New(name string, types int) *Rabbit {
	switch types {
	case 1:
		return newRabbitMq(name, "", "") //simple模式只要队列,会绑定默认的交换机
	case 2:
		return newRabbitMq("", name, "") //simple模式只要队列,会绑定默认的交换机
	}
	return nil
}
