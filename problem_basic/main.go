package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func nearlySimilarRectangles(sides [][]int64) int64 {
	// Write your code here
	counter := int64(0)
	counts := make(map[float64]int64)
	for _, v := range sides {
		if _, ok := counts[float64(v[0])/float64(v[1])]; ok {
			counts[float64(v[0])/float64(v[1])] += 1
		} else {
			counts[float64(v[0])/float64(v[1])] = 0
		}
	}
	for _, v := range counts {
		c := int64(0)
		for i := int64(0); i <= v; i++ {
			c += i
		}
		counter += c
	}
	return counter
}

const fileName = "file.txt"

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(fileName)
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	sidesRows, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	sidesColumns, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var sides [][]int64
	for i := 0; i < int(sidesRows); i++ {
		sidesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var sidesRow []int64
		for _, sidesRowItem := range sidesRowTemp {
			sidesItem, err := strconv.ParseInt(sidesRowItem, 10, 64)
			checkError(err)
			sidesRow = append(sidesRow, sidesItem)
		}

		if len(sidesRow) != int(sidesColumns) {
			panic("Bad input")
		}

		sides = append(sides, sidesRow)
	}

	result := nearlySimilarRectangles(sides)

	fmt.Fprintf(writer, "%d\n", result)

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
