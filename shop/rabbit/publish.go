package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
)

/**
  publish/subscribe
  订阅者模式:消息被路由投递给多个队列,一个消息被多个消费者获取
*/
func Subscribe(Exchange string) *Rabbit {
	//根据不同模式传入不同参数创建实例
	r := newRabbitMq("", Exchange, "")
	//创建连接
	var err error
	r.conn, err = amqp.Dial(r.Url)
	r.Err(err, "连接失败")
	//创建channel
	r.channel, err = r.conn.Channel()
	r.Err(err, "创建管道失败")
	return r
}

/*
 *
 * @param mes
 * @return
 * @author yuxin
 * @creed: Talk is cheap,show me the code
 * @date 2020/3/11 18:03
 */
func (r *Rabbit) PublishSub(mes string) {
	//1.创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout", //确定交换机类型,订阅者模式下为fanout
		true,     //是否持久化
		false,    //是否自动删除
		false,    //true表示这个exchange不可以被client用来推送消息,只用来进行exchange和exchange之间的绑定
		false,
		nil,
	)
	r.Err(err, "创建交换机失败")
	r.channel.Publish(r.Exchange, r.QueueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(mes),
	})
}

/**
 * @description:

 * @return:
 * @author: King
 * @date:
 */
func (r *Rabbit) ConsumeSub() {
	//1.创建交换机
	err := r.channel.ExchangeDeclare(r.Exchange, "fanout", true, false, false, false, nil)
	r.Err(err, "创建交换机失败")
	//2.创建随机队列,因为订阅者模式:消息被路由投递给多个队列,所以需要创建队列,第一个参数设置queue,随机
	q, err := r.channel.QueueDeclare("", false, false, true, false, nil)
	r.Err(err, "创建队列失败")
	//3.将队列绑定到exchange中
	err = r.channel.QueueBind(q.Name, "", r.Exchange, false, nil)
	r.Err(err, "绑定队列失败")
	//4.消费消息
	mes, err := r.channel.Consume(q.Name, "", true, false, false, false, nil)
	ch := make(chan struct{})
	go func() {
		for d := range mes {
			fmt.Println(string(d.Body))
		}
	}()
	<-ch
}
