package main

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"rabbitmq/FanoutExchange/config"
)

func main()  {

	conn, err:= amqp091.Dial(config.RMQADDR)
	failOnError(err,"conn to rabbitmq error")
	defer conn.Close()


	forever := make(chan bool)

	ch, err := conn.Channel()

	failOnError(err,"get channel error")
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
	failOnError(err,"Failed to declare exchange")

	for routine := 0;routine<config.CONSUMERCNT;routine++{
		go func(routineNume int) {
			q,err := ch.QueueDeclare(
				"",
				false,
				false,
				true,
				false,
				nil,
				)
			failOnError(err, "failed to declare a queue")

			err = ch.QueueBind(
				q.Name,
				"",
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

			if err != nil{
				log.Fatal(err)
			}

			for msg := range msgs{
				log.Printf("In %d consume a message: %s\n",routineNume,msg.Body)
			}


		}(routine)
	}

	<- forever
}





func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}