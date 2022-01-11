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
	updOrder1 := pb.Order{
		Id:          "111",
		Items:       []string{"nil","jjj"},
		Description: "init1",
		Price:       55,
		Destination: "before1",
	}
	updOrder2 := pb.Order{
		Id:          "222",
		Items:       []string{"nil","kkk"},
		Description: "init2",
		Price:       55,
		Destination: "before2",
	}
	m := make(map[string]*pb.Order)
	m[updOrder1.Id] = &updOrder1
	m[updOrder2.Id] = &updOrder2
	pb.RegisterOrderManagementServer(s,&server.OrderServer{OrderMap: m})

	log.Printf("Start gRPC listener on port "+ port)

	if err := s.Serve(lis);err != nil{
		log.Fatalf("failed to serve: %v",err)
	}

}
