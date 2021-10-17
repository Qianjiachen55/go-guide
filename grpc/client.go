package main

import (
	"context"
	"fmt"
	"github.com/Qianjiachen55/go-guide/grpc/pb"
	"google.golang.org/grpc"
)

func main()  {
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil{
		fmt.Println("!:",err)
	}
	defer conn.Close()

	pb.NewGreeterClient(conn)

	c := pb.NewGreeterClient(conn)
	r,err := c.SayHello(context.Background(),&pb.HelloRequest{Name: "qwer..."})
	if err != nil{
		fmt.Println("err:,,",err)
	}
	fmt.Println("greeting: ",r.Message)
}
