package listen

import (
	"gRPC_demo/gRPC_DEMO/pb"
	"gRPC_demo/gRPC_DEMO/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func OrderListen()  {
	port := ":50051"
	lis,err := net.Listen("tcp",port)

	if err != nil{
		log.Fatalf("failed to listen: %v",err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s,&server.OrderServer{})

	log.Printf("Start gRPC listener on port "+ port)

	if err := s.Serve(lis);err != nil{
		log.Fatalf("failed to serve: %v",err)
	}

}
