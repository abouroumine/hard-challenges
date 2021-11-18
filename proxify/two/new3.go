package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
)

/*
 * Complete the 'readFile' function below.
 *
 * The function accepts following parameters:
 *  1. int m
 *  2. int n
 *  3. chan []byte bytesChannel
 *  4. chan error errChannel
 */

func readFile(m, n int, bytesChannel chan []byte, errChannel chan error) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		errChannel <- err
	} else {
		b, er := ioutil.ReadFile(filename)
		if er != nil {
			errChannel <- er
			close(bytesChannel)
		} else {
			for i := 0; i+m < len(b); i += m {
				bytesChannel <- b[i : i+n]
			}
			errChannel <- nil
		}
	}
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int(mTemp)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)

	inputString := readLine(reader)

	file, err := os.Create(filename)
	checkError(err)
	defer file.Close()
	_, err = file.WriteString(inputString)
	checkError(err)
	bytesChannel, errChannel := make(chan []byte), make(chan error)
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	allocBefore := ms.Alloc
	go readFile(m, n, bytesChannel, errChannel)
	go func() {
		for {
			select {
			case next := <-bytesChannel:
				fmt.Println(string(next))
			default:
			}
		}
	}()
	err = <-errChannel
	if err == io.EOF {
		runtime.ReadMemStats(&ms)
		allocAfter := ms.Alloc
		if allocAfter-allocBefore > 10000 {
			fmt.Println("Memory usage limit exceeded")
		}
	} else {
		checkError(err)
	}
}

const filename = "output"

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
