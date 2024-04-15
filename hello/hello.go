package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	// "reflect"
)

const delay = 5
const times = 3

func main() {
	introduction()

	for {
		menu()
		command := read()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing logs...")
			showLogs()
		case 0:
			fmt.Println("Exiting")
			os.Exit(0)
		default:
			fmt.Println("Don't know this command")
			os.Exit(-1)
		}
	}
}

func introduction() {
	name := "Pogan"
	version := 1.1

	fmt.Println("Hello Mr.", name)
	fmt.Println("This program is in version", version)
}

func read() int {
	var command int
	fmt.Scan(&command)
	return command
}

func menu() {
	fmt.Println("1 - Initiate Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	urls := readFile()
	for i := 0; i < times; i++ {
		for i, url := range urls {
			fmt.Println("Testing url", i, ":", url)
			testUrl(url)
		}
		time.Sleep(delay * time.Second)
	}
	fmt.Println("")
}

func testUrl(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Url:", url, "loaded successfully!")
		log(url, true)
	} else {
		fmt.Println("Url:", url, "has a problem. Status Code:", resp.StatusCode)
		log(url, false)
	}
}

func readFile() []string {
	var urls []string

	file, err := os.Open("urls.txt")

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	reader := bufio.NewReader(file)
	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)

		if row != "" {
			urls = append(urls, row)
		}

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return urls
}

func log(url string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + url + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func showLogs() {
	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
