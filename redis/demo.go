package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)


func conn()  {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	err := client.Set(ctx,"name","saizige",0).Err()

	defer client.Close()
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("success!")
	client.Close()
	return
}
