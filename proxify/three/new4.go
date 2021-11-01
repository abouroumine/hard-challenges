package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 * Complete the 'RunHttpRequest' function below (and imports if necessary).
 *
 * The function is expected to return a *City and an error.
 * The function accepts following parameters:
 *  1. Message msg
 */

func RunHttpRequest(msg Message) (*City, error) {
	if msg.Token == goodToken && (msg.Index >= 0 && msg.Index < int32(len(store))) {
		return &store[msg.Index], nil
	}
	return nil, errors.New("not authorized")
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	//writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	startTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	start := int(startTemp)

	messagesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var messages []Message

	for i := 0; i < int(messagesCount); i++ {
		messagesItem := readLine(reader)
		var message Message
		err := json.Unmarshal([]byte(messagesItem), &message)
		checkError(err)
		messages = append(messages, message)
	}

	store = store[start:]
	http.HandleFunc("/", rootHandler)
	go http.ListenAndServe(portSuffix, nil)
	time.Sleep(100 * time.Millisecond)
	for _, msg := range messages {
		resp, err := RunHttpRequest(msg)
		if err != nil {
			//fmt.Fprintf(writer, "%s\n", err.Error())
			fmt.Println("Error: ", err.Error())
			continue
		}
		fmt.Println("Name: ", resp.Name)
		fmt.Println("Population: ", resp.Population)
	}
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

type City struct {
	Name       string
	Population int32
}

type Message struct {
	Token string
	Index int32
}

type Request struct {
	Index int32
}

type Response struct {
	Data   City
	ErrMsg string
}

const portSuffix = ":3333"

var address = "http://127.0.0.1" + portSuffix

const goodToken = "goodtoken"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	var resp Response
	if token == goodToken {
		var reqBody Request
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			resp = Response{
				ErrMsg: err.Error(),
			}
		} else {
			resp = Response{
				Data: store[reqBody.Index],
			}
		}
	} else {
		resp = Response{
			ErrMsg: "not authorized",
		}
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

var store = []City{
	{
		Name:       "Warsaw",
		Population: int32(1793579),
	},
	{
		Name:       "Krakow",
		Population: int32(780981),
	},
	{
		Name:       "Lodz",
		Population: int32(677286),
	},
	{
		Name:       "Wroclaw",
		Population: int32(643782),
	},
	{
		Name:       "Poznan",
		Population: int32(533830),
	},
	{
		Name:       "Gdansk",
		Population: int32(471525),
	},
	{
		Name:       "Szczecin",
		Population: int32(400990),
	},
	{
		Name:       "Bydgoszcz",
		Population: int32(346739),
	},
	{
		Name:       "Lublin",
		Population: int32(339547),
	},
	{
		Name:       "Bialystok",
		Population: int32(297585),
	},
	{
		Name:       "Katowice",
		Population: int32(291774),
	},
	{
		Name:       "Gdynia",
		Population: int32(245867),
	},
	{
		Name:       "Czestochowa",
		Population: int32(219278),
	},
	{
		Name:       "Radom",
		Population: int32(210532),
	},
}
