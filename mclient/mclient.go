package main

import (
	"esni-shell-channel/util"
	"io"
	"log"
	"net"
	"os"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"

	"golang.org/x/crypto/ssh"
)

var (
	serverAddress = "www.zhanmabigdata.top"
	//password      = os.Args[2]
	user          = os.Args[1]
	msgPort       = "9658"
	cliport       = "9656"
)

var password string

func handleMsg(conn net.Conn) {

	buffer := make([]byte, 1)
	log.Println("staring to wait message")

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("some problems occurred while handling message: ", err)
			return
		}

		if n != 0 {
			if 1 == int(buffer[0]) {

				log.Println("receiving a create request, staring to create local ssh connection")
				conn, err := net.Dial("tcp", serverAddress+":"+cliport)
				if err != nil {
					log.Println("some problems occurred while creating local ssh connection: ", err)
					return
				}
				defer conn.Close()
				log.Println("successfully created local ssh client")

				log.Println("staring to create local ssh client")
				sshConfig := util.GetSSHConf(user, password)
				cli, err := ssh.Dial("tcp", "localhost:22", sshConfig)
				if err != nil {
					log.Println("some problems occurred while creating local ssh client: ", err)
				}
				defer cli.Close()
				log.Println("successfully created local ssh client")

				log.Println("staring to listen connection from server")
				session, err := cli.Dial("tcp", "localhost:22")
				if err != nil {
					log.Println("some problems occurred while lisenling connection:", err)
				}
				defer session.Close()

				go func() {
					log.Printf("creating channel local to %s", conn.LocalAddr())
					for {
						n, err := io.Copy(session, conn)
						if err != nil {
							log.Println("some error occurred during forwarding", err)
							session.Close()
							conn.Close()
							return

						}
						if n == 0 {
							log.Println("got 0 bytes from connection or local ssh client")
							log.Printf("try to close the channel local to %s", conn.LocalAddr())
							session.Close()
							conn.Close()
							return
						}
					}

				}()

				go func() {
					for {
						log.Printf("creating channel %s to local", conn.LocalAddr())
						n, err := io.Copy(conn, session)
						if err != nil {
							log.Println("some error occurred during forwarding", err)
							session.Close()
							conn.Close()
							return

						}
						if n == 0 {
							log.Println("got 0 bytes from connection or local ssh client")
							log.Printf("try to close the channel %s to local", conn.LocalAddr())
							session.Close()
							conn.Close()
							return
						}
					}
				}()

			}
		}

	}

}

func main() {

	fmt.Println("Enter password: ")
	p, errs := terminal.ReadPassword(0)
	password = string(p)

	if errs != nil {
		fmt.Println("\nCould not read password:")
		log.Fatal(errs)
		os.Exit(1)
	}

	log.Println("staring esni-shell-channel multi client now")

	msgConn, err := net.Dial("tcp", serverAddress+":"+msgPort)
	if err != nil {
		log.Println("error to created message connection: %v", err)
		return
	}
	log.Println("successfully created message connection")
	defer msgConn.Close()

	log.Println("starring to handle message connection")
	go handleMsg(msgConn)

	// lock
	for {
	}

}
