package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	url := "https://golang.org/"

	httpCli := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(http.MethodOptions, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := httpCli.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)
	io.Copy(os.Stdout, resp.Body)
}
