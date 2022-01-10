package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"ggrpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

func main()  {

	cert,_:=tls.LoadX509KeyPair("clientCert/client.pem","clientCert/client.key")
	certPool := x509.NewCertPool()
	ca,_ := ioutil.ReadFile("clientCert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates:                []tls.Certificate{cert},
		ServerName:                  "localhost",
		RootCAs:                   certPool,
	})

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
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
