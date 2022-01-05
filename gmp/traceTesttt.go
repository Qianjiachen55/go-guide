package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	count := 0
	m := sync.Mutex{}
	wg := sync.WaitGroup{}
	//创建文件
	//f, err := os.Create("trace.out")
	//if err != nil{
	//	panic(err)
	//}
	//
	//defer f.Close()
	//
	//err = trace.Start(f)
	//
	//if err != nil{
	//	panic(err)
	//}
	wg.Add(100)
	demo := func() {
		m.Lock()
		//defer m.Unlock()
		count += 1
		m.Unlock()
		time.Sleep(time.Second)
		wg.Done()

	}

	for i:=0;i<100;i++{
		go demo()
	}
	wg.Wait()
	fmt.Println(count)

	//trace.Stop()

}
