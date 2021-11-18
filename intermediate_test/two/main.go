package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'TCPServer' function below.
 *
 * The function accepts chan bool ready as a parameter.
 */

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func TCPServer(ready chan bool) {
	s, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return
	}
	l, err := net.ListenTCP("tcp", s)
	if err != nil {
		return
	}
	ready <- true
	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		buffer := make([]byte, 1024)
		_, err = conn.Read(buffer)
		if err == io.EOF {
			break
		}
		msg := string(buffer)
		reversed := Reverse(msg)
		_, _ = conn.Write([]byte(reversed))
		_ = conn.Close()
	}
}

const maxBufferSize = 1024
const address = "127.0.0.1:3333"

const fileName = "file.txt"

func main() {
	stdout, err := os.Create(fileName)
	checkError(err)

	defer stdout.Close()

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	messagesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var messages []string

	for i := 0; i < int(messagesCount); i++ {
		messagesItem := readLine(reader)
		messages = append(messages, messagesItem)
	}

	ready := make(chan bool)
	go TCPServer(ready)
	<-ready
	reversed, err := tcpClient(messages)
	if err != nil {
		panic(err)
	}
	for _, msg := range reversed {
		fmt.Fprintf(writer, "%s\n", msg)
	}
	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func tcpClient(messages []string) ([]string, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return []string{}, err
	}

	reversed := []string{}

	for _, msg := range messages {

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			return []string{}, err
		}
		_, err = conn.Write([]byte(msg))
		if err != nil {
			return []string{}, err
		}
		reply := make([]byte, maxBufferSize)
		n, err := conn.Read(reply)
		if err != nil {
			return []string{}, err
		}
		fmt.Println("The Message Returned: ", string(reply))
		reversed = append(reversed, string(reply[:n]))
		conn.Close()
	}

	return reversed, nil
}
