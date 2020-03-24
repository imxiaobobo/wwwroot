/*
   @Time : 2020/3/14 12:00
   @Author : wangbo
   @File : topic
*/
package rabbit

type Topic struct {
	rabbit
}

func NewTopic() *Topic {
	r := newRabbitMq("", "", "")
	return &Topic{*r}
}
