package main

import (
	"fmt"
	"net/http"
	"os"
)



func main() {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:7890")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7890")

	url := "https://zh.wikipedia.org/zh/中华人民共和国"

	if response, err := http.Get(url); err == nil {
		fmt.Println(response)
	}else {
		fmt.Println(err)
	}

}
