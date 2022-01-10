package main

import (
	"crypto/tls"
	"crypto/x509"
	"ggrpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"net"
)

func main()  {
	cert,_:=tls.LoadX509KeyPair("serverCert/server.pem","serverCert/server.key")
	certPool := x509.NewCertPool()
	ca,_ := ioutil.ReadFile("serverCert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates:                []tls.Certificate{cert},
		ClientAuth:                  tls.RequireAndVerifyClientCert,
		ClientCAs:                   certPool,
	})

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	service.RegisterProdServiceServer(rpcServer,new(service.ProdService))

	listen, _ := net.Listen("tcp", ":8081")

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	rpcServer.ServeHTTP(writer,request)
	//})
	//
	//httpServer := &http.Server{
	//	Addr: ":8081",
	//	Handler: mux,
	//}

	//httpServer.ListenAndServeTLS()
	rpcServer.Serve(listen)
}
