package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	ADDR = "localhost:50050"
)

func main() {
	conn, err := net.Dial("tcp", ADDR)
	if err != nil {
		fmt.Println("Error Connection: ", err.Error())
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err.Error())
			break
		}
		if text == "stop\n" {
			break
		}
		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error Writing: ", err.Error())
			break
		}
	}
}
