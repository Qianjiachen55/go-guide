package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	var err error = nil
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints:            []string{"localhost:2379"},
		AutoSyncInterval:     0,
		DialTimeout:          5 * time.Second,
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
		MaxCallSendMsgSize:   0,
		MaxCallRecvMsgSize:   0,
		TLS:                  nil,
		Username:             "",
		Password:             "",
		RejectOldCluster:     false,
		DialOptions:          nil,
		Context:              nil,
		Logger:               nil,
		LogConfig:            nil,
		PermitWithoutStream:  false,
	})

	var ctx = context.Background()


	defer etcdCli.Close()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_,err = etcdCli.Put(ctx,"name","qwefqwe")

	if err!= nil{
		fmt.Println("set err: ",err)
		return
	}

	res, err :=etcdCli.Get(ctx,"name")

	if err !=nil{
		fmt.Println("get err : ",err)
	}

	for i,ev := range res.Kvs{
		fmt.Println("res index :",i,string(ev.Key),string( ev.Value))
	}


}
