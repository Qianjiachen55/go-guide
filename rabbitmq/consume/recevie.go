package consume

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"math/rand"
	"rabbitmq/util"
)




func Receive(q amqp091.Queue,target int) {
	// 1. 连接

	conn := util.GetConn()
	defer conn.Close()

	// 2. 创建通道
	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 3. 声明发送的队列，讲消息发布到队列
	//q, err := ch.QueueDeclare(
	//	queueName,
	//	false,
	//	false,
	//	false,
	//	false,
	//	nil,
	//)
	//util.FailOnError(err, "Failed to declare a queue")



	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	util.FailOnError(err, "Failed to declare a consumer")

	forever := make(chan bool)
	//ch.Ack()
	go func() {
		for d := range msgs {
			//log.Printf("%d: Received a message: %s ", target,d.Body)
			//ch.Ack()
			if rand.Int()%3==0 {
				err := d.Ack(false)
				util.FailOnError(err,fmt.Sprintf("failed to ack msg: %s",d.Body))
				fmt.Printf("%d runtine,%s devliver success\n",target,d.Body)
			}else{
				err := d.Nack(false,true)
				util.FailOnError(err,fmt.Sprintf("failed to nack msg: %s",d.Body))
				fmt.Printf("%d runtine ,%s resend to queue\n",target,d.Body)
			}

			//d.Nack()
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
