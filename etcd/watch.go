package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main()  {
	var err error
	var ctx = context.Background()
	cli,err :=clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: time.Second*5,
	})

	if err != nil{
		fmt.Println("connect error: ",err)
		return
	}

	watchCh := cli.Watch(ctx,"name")

	for wresp := range watchCh{
		for _,evt := range wresp.Events{
			fmt.Println(evt)
		}

	}
}