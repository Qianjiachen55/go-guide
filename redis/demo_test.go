package main

import (
	"fmt"
	"testing"
)

func TestConn(t *testing.T) {
	conn()

}

func TestDemo(t *testing.T)  {
	//for true {
	//	svc("1")
	//}
	count := 0
	for true{
		count ++
		fmt.Println(count)
		svc("4")
	}
}
