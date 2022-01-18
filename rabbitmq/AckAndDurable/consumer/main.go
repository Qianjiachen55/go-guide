package main

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"rabbitmq/AckAndDurable/util"
)

//RMQADDR     = "amqp://guest:guest@172.17.84.205:5672/"
//QUEUENAME   = "msgQueueWithPersist"
//PRODUCERCNT = 5


const url = "amqp://guest:guest@localhost:5672/"

func main()  {
	conn, err := amqp091.Dial(url)
	failOnFail(err,"failed to connect to RabbitMq")
	defer conn.Close()
	consumerCOUNT := 20
	queueName := "msgQueueWithPersist"
	forever := make(chan bool)

	for routine := 0;routine<consumerCOUNT;routine++{
		go func(routineNum int) {
			ch,err := conn.Channel()
			failOnFail(err,"failed to open channel")
			defer ch.Close()

			q, err := ch.QueueDeclare(
				queueName,
				true,
				false,
				false,
				false,
				nil,
				)
			failOnFail(err,"failed to declare a queue")

			err = ch.Qos(1,0,false)
			failOnFail(err,"set QOS failed")

			msgs, err := ch.Consume(
				q.Name,
				"MsgConsume",
				false,
				false,
				false,
				false,
				nil,
				)

			failOnFail(err,"consume failed")

			for msg := range msgs{
				log.Printf("In %d consume a message: %s\n",routineNum,msg.Body)
				msg.Ack(false)
			}

		}(routine)
	}

	<- forever

}

func failOnFail(err error,msg string)  {
	util.FailOnError(err,msg)
}