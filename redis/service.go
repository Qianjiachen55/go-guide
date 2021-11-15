package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"math"
	"time"
)

var ctx = context.Background()

func con() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
		PoolSize: 30,
		MinIdleConns: 10,

	})

	return rdb
}

func svc(id string) {

	rdb := con()
	key := "compid" + id
	value := rdb.Get(ctx, key)
	if value.Err() == redis.Nil {
		//not exits
		rdb.SetEX(ctx,key, math.MaxInt64-10, 10*time.Second)
	}else {
		//exits
		res :=rdb.Incr(ctx,key)
		if res.Err()!=nil{
			fmt.Println(res.Err())
			time.Sleep(10*time.Second)
		}
		business(res.Val()-math.MaxInt64+10)
	}

	rdb.Close()
}

func business(val interface{}) {

	fmt.Println("业务执行"+cast.ToString(val))
}
