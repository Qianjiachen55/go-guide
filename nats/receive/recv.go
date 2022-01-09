package receive

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cast"
)

func Recv(cn *nats.Conn,topicName string ) {
	//fmt.Println("recv")
	cn.Subscribe(topicName, func(msg *nats.Msg) {
		//fmt.Println("--------------")
		fmt.Println(cast.ToString(msg.Data))
		//fmt.Println("--------------")

	})
}
