package main

import "fmt"

func main() {
	ch := make(chan int,1)
	for i:=0;i<10;i++{
		select {
		case x := <-ch:
			fmt.Println(x)
			fmt.Println("case1")
		case ch <-i:
			fmt.Println("case2")
		default:
			fmt.Println("nothing")
		}
	}
}
