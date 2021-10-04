package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'ModifyString' function below and add imports if needed.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING str as parameter.
 */

func ModifyString(str string) string {
	s := strings.TrimLeft(str, " ")
	s = strings.TrimRight(s, " ")
	s2 := ""
	fmt.Println(s)
	for v := range s {
		value, err := strconv.Atoi(string(s[v]))
		if err != nil {
			s2 = s2 + string(s[value])
		}
	}
	s3 := ""
	for i := len(s2) - 1; i >= 0; i-- {
		s3 = s3 + string(s2[i])
	}
	return s3
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	str := readLine(reader)

	result := ModifyString(str)

	fmt.Fprintf(writer, "%s\n", result)

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
