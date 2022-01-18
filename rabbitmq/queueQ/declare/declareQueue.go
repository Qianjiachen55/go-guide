package declare

import "github.com/rabbitmq/amqp091-go"

func DeclareQueue(ch *amqp091.Channel,name string, durable, autoDelete, exclusive, noWait bool, args amqp091.Table)(amqp091.Queue,error) {

	return ch.QueueDeclare(name,durable,autoDelete,exclusive,noWait,args)

}
