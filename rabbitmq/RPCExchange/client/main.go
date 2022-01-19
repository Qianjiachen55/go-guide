package main

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"math/rand"
	"os"
	config "rabbitmq/RPCExchange/conf"
	"strconv"
)

func main() {

	conn, err := amqp091.Dial(config.RMQADDR)
	failOnError(err, "failed to connect to rabbitmq")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	var msg string

	if len(os.Args) < 2 {
		msg = strconv.Itoa(rand.Int())
	} else {
		msg = os.Args[1]
	}

	resQueue, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "failed to declare a response queue")

	correlationID := strconv.Itoa(rand.Intn(10))

	err = ch.Publish(
		"",
		config.QUEUENAME,
		false,
		false,
		amqp091.Publishing{
			ContentType:   "text/plain",
			CorrelationId: correlationID,
			ReplyTo:       resQueue.Name,
			Body:          []byte(msg),
		},
	)

	log.Printf(" [x] Sent %s", msg)
	failOnError(err, "failed to publist a message")

	respMsgs, err := ch.Consume(
		resQueue.Name,
		"",
		true,
		true,
		false,
		false,
		nil,
	)

	for item := range respMsgs{
		if item.CorrelationId ==correlationID{
			fmt.Println("response: ",string(item.Body))
			break
		}
	}

}

func failOnError(err error, msg interface{}) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
