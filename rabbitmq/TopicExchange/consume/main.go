package main

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"rabbitmq/TopicExchange/config"
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
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "failed to declare exchange")

	topics := []string{"#", "info.payment.*", "*.log", "debug.payment.#"}

	for cnt := 0; cnt < len(topics); cnt++ {
		go func(routingNum int) {
			q, err := ch.QueueDeclare(
				"",
				false,
				false,
				true,
				false,
				nil,
			)
			failOnError(err,"Failed to declare a queue")

			err = ch.QueueBind(
				q.Name,
				topics[routingNum],
				config.EXCHANGENAME,
				false,
				nil,
				)
			failOnError(err,"failed to bind exchange")

			msgs,err := ch.Consume(
				q.Name,
				"",
				true,
				false,
				false,
				false,
				nil,
				)
			failOnError(err,"Failed to register a consumer")

			for msg := range msgs{
				log.Printf("In %s consume a message :%s\n",topics[routingNum],msg.Body)
			}

		}(cnt)
	}
	<- forever
}

func failOnError(err error, msg interface{}) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
