package util

import (
	"github.com/nats-io/nats.go"
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetConn()  (*nats.Conn , error){
	nc, err := nats.Connect(nats.DefaultURL)
	return nc,err
}
