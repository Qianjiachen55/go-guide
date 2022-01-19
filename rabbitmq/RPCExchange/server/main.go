package main

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	config "rabbitmq/RPCExchange/conf"
	"time"
)

func main() {
	conn, err := amqp091.Dial(config.RMQADDR)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.Qos(config.SERVERINSTANCESCNT, 0, false)

	forever := make(chan bool)


	for routine := 0; routine < config.SERVERINSTANCESCNT; routine++ {
		go func(routineNum int) {
			q, err := ch.QueueDeclare(
				config.QUEUENAME,
				false,
				false,
				false,
				false,
				nil,
			)

			failOnError(err, "failed to declare a queue")

			msgs, err := ch.Consume(
				q.Name,
				"",
				false,
				false,
				false,
				false,
				nil,
			)

			for msg := range msgs {
				log.Printf("In %d start consuming message: %s\n", routineNum, msg.Body)

				bookName := "queried: " + string(msg.Body)
				time.Sleep(time.Second / 2)

				err = ch.Publish(
					"",
					msg.ReplyTo,
					false,
					false,
					amqp091.Publishing{
						ContentType:   "text/plain",
						CorrelationId: msg.CorrelationId,
						Body:          []byte(bookName),
					},
				)

				msg.Ack(false)
			}

		}(routine)

	}

	<- forever

}

func failOnError(err error, msg interface{}) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
