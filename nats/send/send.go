package send

import "github.com/nats-io/nats.go"

func Send(nc *nats.Conn,topicName string,data []byte)  {
	nc.Publish(topicName,data)
}
