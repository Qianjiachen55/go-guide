package client

import (
	"context"
	"fmt"
	"gRPC_demo/gRPC_DEMO/pb"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"io"
	"log"
)

func GetAddOrder()  {

	address := "localhost:50051"
	conn,err := grpc.Dial(address,grpc.WithInsecure())
	if err != nil{
		log.Fatalf("did not connect: %v",err)
	}
	defer conn.Close()
	id := uuid.New().String()
	orderMgtClient := pb.NewOrderManagementClient(conn)
	res,err := orderMgtClient.AddOrder(context.Background(),&pb.Order{
		Id:          id,
		Items:       []string{"Google","phone"},
		Description: "an apple phone",
		Price:       100,
		Destination: "asdfasdfa",
	})
	if err != nil{
		fmt.Println("add err: ",err)
	}else {
		fmt.Println("add success id: ",res.Value)
	}


	retrievedOrder , err := orderMgtClient.GetOrder(context.Background(),&wrappers.StringValue{Value: id})

	log.Println("GetOrder Response -> : ",retrievedOrder)
}

func SearchOrders()  {
	ctx := context.Background()
	address := "localhost:50051"
	conn,err := grpc.Dial(address,grpc.WithInsecure())
	if err != nil{
		log.Fatalf("did not connect: %v",err)
	}
	defer conn.Close()

	c := pb.NewOrderManagementClient(conn)

	searchStream, _ := c.SearchOrders(ctx,&wrappers.StringValue{Value: "Google"})

	for {
		searchOrder,err := searchStream.Recv()
		if err == io.EOF{
			break
		}

		log.Print("Search Result : ",searchOrder)
	}
	fmt.Println("--------")
}