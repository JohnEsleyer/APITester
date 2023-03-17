package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli-app [GET/POST] [URL] [optional:POST_DATA]")
		return
	}

	method := os.Args[1]
	url := os.Args[2]
	var postBody []byte

	if method == "POST" && len(os.Args) >= 4 {
		postBody = []byte(os.Args[3])
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Status code:", resp.StatusCode)
	fmt.Println("Headers:")
	for key, values := range resp.Header {
		fmt.Printf("%s: %s\n", key, values)
	}
	fmt.Println("Body:")
	fmt.Println(string(body))
}
