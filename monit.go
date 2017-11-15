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
)

const delay = 3

func main() {
	startMonitoring()
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	hosts := readFile()

	for {
		for _, h := range hosts {
			fmt.Println("Check", h)
			checkHost(h)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func checkHost(host string) {
	resp, err := http.Get(host)

	if err != nil {
		fmt.Println("ERROR:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println(host, "OK")
		recordLogs(host, true)
	} else {
		fmt.Println(host, "ERROR", resp.StatusCode)
		recordLogs(host, false)
	}
}

func readFile() []string {
	var hosts []string
	file, err := os.Open("hosts.txt")

	if err != nil {
		fmt.Println("ERROR:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		hosts = append(hosts, line)
		if err == io.EOF {
			break
		}
	}

	file.Close()
	return hosts
}

func recordLogs(host string, status bool) {
	file, err := os.OpenFile("status.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("ERRO:", err)
	}

	file.WriteString(time.Now().Format("Jan 02 15:04:05") + " - " + host + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}
