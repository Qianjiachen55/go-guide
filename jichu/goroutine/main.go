package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func hello(i int)  {

	fmt.Println("hello",i)
	wg.Done()
}

func main(){

	wg.Add(1000)
	for i:=0;i<1000;i++{
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}


	fmt.Println("main")

	wg.Wait()
}
