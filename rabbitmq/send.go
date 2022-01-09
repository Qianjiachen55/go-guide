package main

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 1. 连接

	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 2. 创建通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 3. 声明发送的队列，讲消息发布到队列
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	go func() {
		for {
			t := time.NewTimer(time.Second/2)
			select {
			case <-t.C:
				publish(ch, q)
				log.Println("send success")
			}
		}
	}()

	select {

	}

}

func publish(ch *amqp091.Channel, q amqp091.Queue) {
	time.Location{}
	body := "Hello World! " + time.Now().Format("2006-01-02 15:04:05")
	err := ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "publish failed")
}
