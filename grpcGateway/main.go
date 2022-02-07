package main

import (
	"context"
	pb "gateway/proto/helloWorld"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)


type server struct {
	pb.UnimplementedGreeterServer
}

func (s server) SayHello(ctx context.Context,in *pb.HelloRequest)(*pb.HelloReply,error) {
	return &pb.HelloReply{Message: "hello :" + in.Name},nil
}

func main()  {
	s := grpc.NewServer()

	pb.RegisterGreeterServer(s,&server{})

	lis,err := net.Listen("tcp",":50051")
	if err != nil{
		log.Fatalf("failed to listen: %v",err)
	}

	go func() {
		if err := s.Serve(lis);err != nil{
			log.Fatalf("failed to server: %v",err)
		}
	}()

	conn,err := grpc.Dial(":50051",grpc.WithInsecure(),)

	if err != nil{
		log.Fatalln("failed to dial server:",err)
	}

	gwmux := runtime.NewServeMux()

	err = pb.RegisterGreeterHandler(context.Background(),gwmux,conn)
	if err != nil{
		log.Fatalln("Failed to register gateway:", err)

	}

	gwServer := &http.Server{
		Addr:              "0.0.0.0:8081",
		Handler:           gwmux,
	}
	gwServer.ListenAndServe()

}
