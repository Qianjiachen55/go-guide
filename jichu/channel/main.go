package main

import (
	"fmt"
)

func main()  {
	var ch1 chan int
	ch1 = make(chan int,1)
	ch1 <- 10
	x := <- ch1
	fmt.Println(x)
	close(ch1)
}
