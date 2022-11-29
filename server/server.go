package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func forward(conn1 net.Conn, conn2 net.Conn) {
	fmt.Println("create channel")
	for {
		n, err := io.Copy(conn1, conn2)
		if err != nil {
			fmt.Println("server forward error")
		} else {
			fmt.Printf("server forward %d bytes data", n)
		}

	}
}

func main() {

	listener, err := net.Listen("tcp", "ESNI-Master:"+os.Args[1])
	if err != nil {
		fmt.Println("listener error: ", err)
	}
	defer listener.Close()

	i := 1
	var conn1 net.Conn
	var conn2 net.Conn
	for {
		if i == 1 {
			cur, err := listener.Accept()
			conn1 = cur
			if err != nil {
				fmt.Println("connect error: ", err)
			}

			fmt.Println("connect success")
			i++
			defer conn1.Close()
		}
		if i == 2 {
			cur, err := listener.Accept()
			conn2 = cur
			if err != nil {
				fmt.Println("connect error: ", err)
			}
			fmt.Println("cconnect success")
			go forward(conn1, conn2)
			go forward(conn2, conn1)
			i++
			defer conn2.Close()
		}
	}

}
