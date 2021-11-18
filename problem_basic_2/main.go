package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func maxCost(cost []int32, labels []string, dailyCount int32) int32 {
	// Write your code here
	maxcost := int32(0)
	counter := int32(0)
	currentMax := int32(0)
	for i := 0; i < len(cost); i++ {
		v := labels[i]
		c := cost[i]
		if v == "legal" {
			currentMax += c
			counter += 1
		} else {
			currentMax += c
		}
		if dailyCount == counter {
			if maxcost < currentMax {
				maxcost = currentMax
			}
			currentMax = 0
			counter = 0
		}
	}
	return maxcost
}

const fileName = "file.txt"

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(fileName)
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	costCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var cost []int32

	for i := 0; i < int(costCount); i++ {
		costItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		costItem := int32(costItemTemp)
		cost = append(cost, costItem)
	}

	labelsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var labels []string

	for i := 0; i < int(labelsCount); i++ {
		labelsItem := readLine(reader)
		labels = append(labels, labelsItem)
	}

	dailyCountTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	dailyCount := int32(dailyCountTemp)

	result := maxCost(cost, labels, dailyCount)

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
