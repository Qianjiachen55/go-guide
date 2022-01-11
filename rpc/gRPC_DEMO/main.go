package main

import (
	"fmt"
	"gRPC_demo/gRPC_DEMO/global"
	"gRPC_demo/gRPC_DEMO/listen"
	"github.com/spf13/cast"

	//"fmt"
	"gRPC_demo/gRPC_DEMO/client"
	//"github.com/spf13/cast"
	"time"
)

func main()  {

	go listen.ProductListen()

	//layout := "2006-01-02 15:04:05"

	for {
		t := time.NewTimer(time.Second)
		select {
		case <-t.C:

			fmt.Println(cast.ToString(time.Now().Format(global.LAYOUT)))

			client.Send()

		}
	}

}
