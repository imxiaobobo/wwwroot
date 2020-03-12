package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
)

//订阅模式结构体
type Publish struct {
	rabbit
}

/**
  创建订阅者模式实例
*/
func NewPub(Exchange string) *Publish {
	rabbit := newRabbitMq("", Exchange, "")
	p := &Publish{*rabbit}
	return p
}

func (p *Publish) Publish(msg string) {
	/**
	  创建交换机
	  name 交换机名
	  kind 确定交换机名,订阅模式下为fanout
	  durable 是否持久化
	  autoDelete 是否自动删除
	  internal true表示这个exchange不可以被client用来推送消息,只用来进行exchange和exchange之间的绑定
	  noWait 是否阻塞
	  args
	*/
	err := p.channel.ExchangeDeclare(p.exchange, "fanout", true, false, false, false, nil)
	Err(err, "创建交换机失败")

	//发送消息 同simple模式
	p.channel.Publish(p.exchange, p.queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	})
}

func (p *Publish) Consume() {
	//创建交换机
	err := p.channel.ExchangeDeclare(p.exchange, "fanout", true, false, false, false, nil)
	Err(err, "创建交换机失败")
	//创建随机队列,因为订阅者模式,消息被路由投递给多个队列,所以需要创建队列,由于随机队列,所以为空
	queue, err := p.channel.QueueDeclare("", false, false, true, false, nil)
	Err(err, "创建队列失败")
	//将队列绑定到交换机
	err = p.channel.QueueBind(queue.Name, "", p.exchange, false, nil)
	Err(err, "绑定队列失败")
	//消费消息
	msg, err := p.channel.Consume(queue.Name, "", true, false, false, false, nil)
	Err(err, "获取消息失败")
	ch := make(chan struct{})
	go func() {
		for d := range msg {
			fmt.Println(string(d.Body))
		}
	}()
	<-ch
}
