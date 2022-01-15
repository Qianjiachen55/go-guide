package main

import (
	"github.com/rabbitmq/amqp091-go"
	"rabbitmq/consume"
	"rabbitmq/declare"
	"rabbitmq/send"
	"rabbitmq/util"
)

func main()  {

	// 1. 连接

	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")


	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 2. 创建通道
	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	queueName := "hello"
	durable := true
	autoDelete := false
	exclusive := false
	noWait := false

	//msg := fmt.Sprintf("{time:%s ,msg:hello}",time.Now().Format(util.LAYOUT))
	msg := "hello"
	queue, err :=declare.DeclareQueue(ch,queueName,durable,autoDelete,exclusive,noWait,nil)



	//go consume.Receive(queueName)



	go send.Send(ch,queue,queueName,msg)
	go send.Send(ch,queue,queueName,msg)
	go send.Send(ch,queue,queueName,msg)
	go send.Send(ch,queue,queueName,msg)
	go send.Send(ch,queue,queueName,msg)
	go send.Send(ch,queue,queueName,msg)
	go send.Send(ch,queue,queueName,msg)
	go send.Send(ch,queue,queueName,msg)


	go consume.Receive(queue,1)
	go consume.Receive(queue,2)
	go consume.Receive(queue,3)

	select {

	}


}
