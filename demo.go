package main

import (
	"fmt"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	fmt.Println("Enter your pwd")
	password, _ := terminal.ReadPassword(0)
	fmt.Println(password)

}
