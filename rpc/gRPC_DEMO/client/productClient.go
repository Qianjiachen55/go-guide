package client

import (
	"context"
	"gRPC_demo/gRPC_DEMO/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Send()  {
	address := "localhost:50051"
	conn,err := grpc.Dial(address,grpc.WithInsecure())
	if err != nil{
		log.Fatalf("did not connect: %v",err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

	name := "Apple iPhone 11"
	description := "hhhhhhhhh"
	ctx, cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	r,err := c.AddProduct(ctx,&pb.Product{
		Name:        name,
		Description: description,
	})

	if err != nil{
		log.Fatalf("could not add product: %v",err)
	}

	log.Printf("send Product ID: %s added successfully",r.Value)
	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil{
		log.Fatalf("Could not get product: %v",err)
	}
	log.Printf("get Product: %s", product.String())

}
