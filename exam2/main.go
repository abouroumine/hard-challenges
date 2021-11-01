package main

import (
	"fmt"
	"net"
)

const (
	ADDR = "localhost:50050"
)

func main() {
	l, err := net.Dial("tcp", ADDR)
	if err != nil {
		fmt.Println("Error 1: ", err.Error())
		return
	}
	defer l.Close()

	msg := "Hello Ayoub\n"

	fmt.Fprintf(l, msg)
}
