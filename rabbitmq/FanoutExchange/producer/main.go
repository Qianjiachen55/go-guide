package main

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"rabbitmq/FanoutExchange/config"
)

func main() {
	conn, err := amqp091.Dial(config.RMQADDR)
	failOnError(err, "fail to conn rabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "fail to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		config.EXCHANGENAME,
		amqp091.ExchangeFanout,
		true,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "failed to declare exchange!!!")
	msgCount := 10

	for count := 0; count < msgCount; count++ {
		msgBody := fmt.Sprintf("hello %d", count)

		ch.Publish(
			config.EXCHANGENAME,
			"",
			false,
			false,
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msgBody),
			},
		)
		log.Printf("[x] Sent %s",msgBody)

	}

	failOnError(err,"failed to publish a message")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
