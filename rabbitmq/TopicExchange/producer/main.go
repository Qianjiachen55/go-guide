package main

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"rabbitmq/TopicExchange/config"
	"strconv"
)

func main() {
	conn, err := amqp091.Dial(config.RMQADDR)
	failOnError(err, "failed to conn rabbitmq!")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		config.EXCHANGENAME,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare exchange")

	routingKey := "info.payment.googlepay"

	msgNum := 10

	for cnt := 0; cnt < msgNum; cnt++ {
		msgBody := "topic: " + strconv.Itoa(cnt)
		err = ch.Publish(
			config.EXCHANGENAME,
			routingKey,
			false,
			false,
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msgBody),
			},
		)
		log.Printf("[x] sent %s",msgBody)
	}
	failOnError(err,"Failed to publish a message")


}

func failOnError(err error, msg interface{}) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
