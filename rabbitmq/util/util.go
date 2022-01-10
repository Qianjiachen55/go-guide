package util

import (
	"github.com/rabbitmq/amqp091-go"
	"log"

)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetConn() *amqp091.Connection {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

