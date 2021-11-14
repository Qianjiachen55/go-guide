package main

import "fmt"

func loop(s *int) {
	for i := 0; i < 100; i++ {
		*s = *s + 1
		fmt.Println(*s)
	}

}

func main() {
	//var a *int
	//b := 0
	//a = &b
	//
	//go loop(a)
	//
	//loop(a)
	//
	//
	//fmt.Println("b = ",b)

	data := make(chan int, 3)
	canQuit := make(chan bool) //阻塞主进程，防止未处理完就退出

	go func() {
		for d := range data {//如果data的缓冲区为空，这个协程会一直阻塞，除非被channel被close
			fmt.Println(d)
		}
		canQuit <- true
	}()

	data <- 5
	data <- 4
	data <- 3
	data <- 2
	data <- 1
	close(data) //用完需要关闭，否则goroutine会被死锁
	<-canQuit //解除阻塞
}