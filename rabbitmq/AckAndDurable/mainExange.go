package main

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	util2 "rabbitmq/queueQ/util"
)

func main()  {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	util2.FailOnError(err,"get conn error")

	ch,err := conn.Channel()
	util2.FailOnError(err,"get channel error")

	exchangeName := "logs"

	//交换机
	err = ch.ExchangeDeclare(
		exchangeName,
		amqp091.ExchangeFanout,
		true,
		false,
		false,
		false,
		nil,
		)
	util2.FailOnError(err,"exchange declare fail")

	//队列
	queue,err :=ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
		)
	util2.FailOnError(err,"Declare queue fail")

	ch.QueueBind(
		queue.Name,
		"",
		exchangeName,
		false,
		nil,
		)
	fmt.Println("等待接受消息")

	//消费者取消息时回调接口


}
