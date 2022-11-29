package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func forward(conn1 net.Conn, conn2 net.Conn) {
	fmt.Printf("create channel forward %s to %s \n", conn1.LocalAddr(), conn2.LocalAddr())
	for {
		n, err := io.Copy(conn1, conn2)
		if err != nil {
			fmt.Println("server forward error: ", err)
			fmt.Printf("try to close channel forward %s to %s \n", conn1.LocalAddr(), conn2.LocalAddr())
			conn1.Close()
			conn2.Close()
			return
		}
		if n == 0 {
			fmt.Println("server forward got 0 bytes: ", err)
			fmt.Printf("try to close channel forward %s to %s \n", conn1.LocalAddr(), conn2.LocalAddr())
			conn1.Close()
			conn2.Close()
			return
		}

	}
}

func handleSession(mconn net.Conn, sListener net.Listener, cListener net.Listener) {

	for {
		sccon, err := sListener.Accept()
		if err != nil {
			fmt.Println("got some error when listen server connection: ", err)
			continue
		}

		// send message
		n, err := mconn.Write([]byte{byte(1)})
		if err != nil || n == 0 {
			fmt.Println("got some error when send message: ", err)
			continue
		}

		// create connection with client
		time.Sleep(1)
		conc, err := cListener.Accept()
		if err != nil {
			fmt.Println("create connection with client error: ", err)
			continue
		}

		// start forward
		go forward(sccon, conc)
		go forward(conc, sccon)

	}

}

// 9656 9657 9658
func main() {

	msgListener, err := net.Listen("tcp", "ESNI-Master:9658")
	if err != nil {
		log.Println("Fail to create message connection listener: ", err)
		return
	}
	defer msgListener.Close()
	fmt.Println("Establish message listener successfully")

	fmt.Println("Wating for first response from message connection")
	msgConn, err := msgListener.Accept()
	if err != nil {
		fmt.Println("Fail to establish message connection: ", err)
		return
	}
	defer msgConn.Close()

	fmt.Println("Establish session listener successfully")
	sessionListener, err := net.Listen("tcp", "ESNI-Master:9657")
	if err != nil {
		fmt.Println("create session listener error: ", err)
		return
	}
	defer sessionListener.Close()

	fmt.Println("create cli listener success")
	clilListener, err := net.Listen("tcp", "ESNI-Master:9656")
	if err != nil {
		fmt.Println("create client listener error: ", err)
		return
	}
	defer clilListener.Close()

	go handleSession(msgConn, sessionListener, clilListener)

	// lock
	for {
	}

}
