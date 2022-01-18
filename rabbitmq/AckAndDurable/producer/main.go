package main

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"rabbitmq/AckAndDurable/util"
	"sync"
)

const url = "amqp://guest:guest@localhost:5672/"

func main() {
	conn, err := amqp091.Dial(url)
	failOnFail(err, "failed to conn rabbitMq")
	defer conn.Close()

	producerCOUNT := 20
	queueName := "msgQueueWithPersist"
	wg := sync.WaitGroup{}
	wg.Add(producerCOUNT)

	for routine := 0; routine < producerCOUNT; routine++ {
		go func(routineNum int) {
			ch, err := conn.Channel()
			failOnFail(err, "failed to open a channel")
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

			for i:=0;i<500;i++{
				msgBody := fmt.Sprintf("msg_%d_%d",routineNum,i)
				err := ch.Publish(
					"",
					q.Name,
					false,
					false,
					amqp091.Publishing{
						DeliveryMode: amqp091.Persistent,
						ContentType: "text/plain",
						Body: []byte(msgBody),
					},
				)
				log.Printf("[x] Sent %s",msgBody)
				failOnFail(err,"failed to publish a message")
			}

			wg.Done()
		}(routine)

	}
	wg.Wait()

	log.Println("All message sent!!!!")

}

func failOnFail(err error, msg string) {
	util.FailOnError(err, msg)
}
