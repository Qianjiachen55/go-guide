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

func GetAddOrder() {

	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	id := uuid.New().String()
	orderMgtClient := pb.NewOrderManagementClient(conn)
	res, err := orderMgtClient.AddOrder(context.Background(), &pb.Order{
		Id:          id,
		Items:       []string{"Google", "phone"},
		Description: "an apple phone",
		Price:       100,
		Destination: "asdfasdfa",
	})
	if err != nil {
		fmt.Println("add err: ", err)
	} else {
		fmt.Println("add success id: ", res.Value)
	}

	retrievedOrder, err := orderMgtClient.GetOrder(context.Background(), &wrappers.StringValue{Value: id})

	log.Println("GetOrder Response -> : ", retrievedOrder)
}

func SearchOrders() {
	ctx := context.Background()
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewOrderManagementClient(conn)

	searchStream, _ := c.SearchOrders(ctx, &wrappers.StringValue{Value: "jjj"})

	for {
		searchOrder, err := searchStream.Recv()
		if err == io.EOF {
			break
		}

		log.Print("Search Result : ", searchOrder)
	}
	fmt.Println("--------")
}

func UpdateOrder() {
	ctx := context.Background()
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewOrderManagementClient(conn)
	updateStream, err := c.UpdateOrders(ctx)

	if err != nil {
		log.Fatalf("%v.UpdateOrders(_) = _, %v", c, err)
	}
	updOrder1 := pb.Order{
		Id:          "111",
		Items:       []string{"nil", "jjj"},
		Description: "111",
		Price:       55,
		Destination: "qwfe",
	}

	if err := updateStream.Send(&updOrder1); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder1, err)
	}

	updOrder2 := pb.Order{
		Id:          "222",
		Items:       []string{"nil", "kkk"},
		Description: "222",
		Price:       222,
		Destination: "222",
	}

	if err := updateStream.Send(&updOrder2); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder2, err)
	}

	updateRes, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", updateStream, err, nil)
	}
	log.Printf("Update orders Res:%s", updateRes)
	log.Println("DONE")

}

func ProcessOrder() {
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	ctx := context.Background()
	c := pb.NewOrderManagementClient(conn)
	streamProcOrder, _ := c.ProcessOrders(ctx)
	if err = streamProcOrder.Send(&wrappers.StringValue{Value: "111"}); err != nil {
		log.Fatalf("%v.send(%v) %v", c, "111", err)
	}

	channel := make(chan struct{})
	go asncClientBidirectionalRPC(streamProcOrder,channel)
	<-channel
}

func asncClientBidirectionalRPC(streamProcOrder pb.OrderManagement_ProcessOrdersClient,c chan struct{})  {
	for{
		combinedShipment, errProcOrder := streamProcOrder.Recv()
		if errProcOrder == io.EOF{
			break
		}
		log.Print("Combined shipment: ",combinedShipment.OrdersList)

	}
	<- c
}
