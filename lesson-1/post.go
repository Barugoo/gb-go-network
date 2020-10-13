package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	resp, err := http.Post(
		"https://postman-echo.com/post",
		"application/json",
		strings.NewReader(`{"foo1":"bar1","foo2":"bar2"}`),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
