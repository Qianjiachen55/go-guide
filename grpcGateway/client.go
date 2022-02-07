package main

import (
	"context"
	pb "gateway/proto/helloWorld"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
	defaultName = "world"
)

func main()  {
	conn,err:=grpc.Dial(address,grpc.WithInsecure(),grpc.WithBlock())

	if err != nil{
		panic(err)
	}

	defer conn.Close()

	c := NewGreeterClient(conn)

	r,err := c.SayHello(context.Background(),&pb.HelloRequest{Name: defaultName})


	if err != nil{
		log.Fatalf("could not greet: %v",err)
	}
	log.Printf("greeting: %s",r.Message)
}