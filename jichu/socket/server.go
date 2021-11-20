package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn)  {
	defer conn.Close()

	for{
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n,err := reader.Read(buf[:])
		if err!=nil{
			fmt.Println(err)
			break
		}
		recv :=string(buf[:n])
		fmt.Println(recv)
		conn.Write([]byte("ok"))
	}
}

func main()  {
	listen,err := net.Listen("tcp","127.0.0.1:20000")
	if err != nil{
		fmt.Println(err)
		return
	}
	for  {
		conn,err := listen.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}
		go process(conn)
	}

}
