package main

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"rabbitmq/DirectExchange/config"
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
		amqp091.ExchangeDirect,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare exchange")

	msgNum := 10
	for cnt := 0; cnt < msgNum; cnt++ {
		msgBody := "hello:" + config.RoutingKeys[cnt%4]
		err = ch.Publish(
			config.EXCHANGENAME,
			config.RoutingKeys[cnt%4],
			false,
			false,
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msgBody),
			},
		)
		log.Printf(" [x] sent %s",msgBody)
	}
	failOnError(err,"failed to publish a message")


}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
