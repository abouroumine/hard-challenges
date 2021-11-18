package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	ADDR     = ":50050"
	fileName = "file.txt"
)

func main() {
	l, err := net.Listen("tcp", ADDR)
	if err != nil {
		fmt.Println("Error Listening: ", err.Error())
		return
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error Connecting: ", err.Error())
		return
	}
	defer conn.Close()
	_, err = os.Create(fileName)
	if err != nil {
		fmt.Println("Error Creating File: ", err.Error())
		return
	}
	data, errChan, done := make(chan string), make(chan error), make(chan bool)
	go writeToFile(done, data, errChan)
	for {
		d, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error Reading: ", err.Error())
			break
		}
		data <- d
	}
	done <- true
	if <-errChan != nil {
		fmt.Println("Failed To Write To File")
	}
}

func writeToFile(done chan bool, data chan string, err chan error) {
	var content string
	for {
		select {
		case <-done:
			err <- os.WriteFile(fileName, []byte(content), 0)
			return
		case next := <-data:
			content += next
		}
	}
}
