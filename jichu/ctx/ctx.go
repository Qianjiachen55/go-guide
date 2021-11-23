package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context)  {
	defer wg.Done()
	LABLE:
	for {
		fmt.Println("woker...")
		time.Sleep(time.Second)
		select {
		case <- ctx.Done():
			break LABLE
		default:

		}
	}
}

func main()  {
	ctx,cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second*3)
	cancel()
	wg.Wait()
	fmt.Println("over!")

}
