package send

import (
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"log"
	util2 "rabbitmq/queueQ/util"
	"sync"
	"time"
)


type Msg struct {
	Id int `json:"ID"`
	Message interface{} `json:"msg"`
	Time string `json:"time"`


}
type count struct {
	Id int
	m sync.Mutex
}

//var m sync.Mutex
//var id = 0

var ID = count{
	Id: 0,
	m:  sync.Mutex{},
}

func Send(ch *amqp091.Channel,q amqp091.Queue, queueName string,msg string) {

	go func() {
		for {
			t := time.NewTimer(time.Second/2)
			select {
			case <-t.C:

				newMsg := Msg{
					Message: msg,
					Time:   time.Now().Format(util2.LAYOUT),
				}
				ID.m.Lock()
				ID.Id++
				//time.Sleep(time.Second/2)

				newMsg.Id = ID.Id
				ID.m.Unlock()
				//time.Sleep(time.Second/2)

				jsonMsg,err := json.Marshal(newMsg)
				//fmt.Println(string(jsonMsg))
				util2.FailOnError(err,"Marshal Failed!")
				publish(ch, q,jsonMsg)
				log.Printf("send success %d", newMsg.Id)
				//newMsg.m.Unlock()
			}
		}
	}()

	select {

	}

}






func publish(ch *amqp091.Channel, q amqp091.Queue,msg []byte) {

	err := ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			DeliveryMode: amqp091.Persistent,
			Body:        msg,
		},
	)
	//ch.
	util2.FailOnError(err, "publish failed")
}

