package main

import (
	"fmt"
	"github.com/spf13/cast"
	"nats/receive"
	"nats/send"
	"nats/util"
	"time"
)

func main() {
	//nc,err := nats.Connect(nats.DefaultURL)
	////nc,err := nats.Connect("nats://127.0.0.1:42")
	//if err != nil{
	//	fmt.Println(err)
	//}
	//
	//err =nc.Publish("foo",[]byte("hello world"))
	//if err != nil{
	//	fmt.Println(err)
	//}
	//
	//_,err =nc.Subscribe("foo", func(msg *nats.Msg) {
	//	fmt.Println("Received a msg: %s\n",cast.ToString(msg.Data))
	//})
	//if err != nil{
	//	fmt.Println(err)
	//}

	// Connect to a server
	nc, err := util.GetConn()
	if err != nil {
		fmt.Println(err)
	}
	// Simple Publisher
	//send.Send(nc,"foo",[]byte("hello world"))
	go func() {
		for {
			t := time.NewTimer(1 * time.Second)
			select {
			case <-t.C:
				send.Send(nc, "foo", []byte("hello world "+cast.ToString(time.Now().Format("2006-01-02 15:04:05"))))
			}
		}

	}()

	//nc.Publish("foo", []byte("Hello World"))
	receive.Recv(nc, "foo")
	// Simple Async Subscriber
	//go func() {
	//	for{
	//		t := time.NewTimer(1*time.Second)
	//		select {
	//		case <-t.C:
	//			receive.Recv(nc,"foo")
	//		}
	//	}
	//}()

	for {
		t := time.NewTimer(1 * time.Second)
		select {
		case <-t.C:
			//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
	}

	nc.Close()
}
