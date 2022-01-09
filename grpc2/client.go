package main

import (
	"context"
	"fmt"
	"ggrpc/service"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()

	client := service.NewProdServiceClient(conn)

	prodRes, err := client.GetProdStock(context.Background(), &service.ProdRequest{ProdId: 12})
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(prodRes.ProdStock)

}
