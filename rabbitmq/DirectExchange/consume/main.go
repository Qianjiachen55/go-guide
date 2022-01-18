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

	forever := make(chan bool)

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

	for routing := 0; routing <= config.CONSUMERCNT; routing++ {
		go func(routingName int) {
			q, err := ch.QueueDeclare(
				"",
				false,
				false,
				true,
				false,
				nil,
			)

			failOnError(err, "failed to declare queue")

			err = ch.QueueBind(
				q.Name,
				config.RoutingKeys[routingName],
				config.EXCHANGENAME,
				false,
				nil,
			)
			failOnError(err, "failed to bind exchange")

			msgs, err := ch.Consume(
				q.Name,
				"",
				true,
				false,
				false,
				false,
				nil,
			)
			failOnError(err,"error to consume a message")

			for msg := range msgs{
				log.Printf("In %s consume a message:%s\n",config.RoutingKeys[routingName],msg.Body)
			}

		}(routing)
	}

	<-forever
}

func failOnError(err error, msg interface{}) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
