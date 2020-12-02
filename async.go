package main

import (
	"fmt"
	"net/http"
)

func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down !!!")
		return
	}
	fmt.Println(url, "is up and running.")
}

func main() {

	urls := []string{
		"https://detik.com/",
		"https://xl.co.id/",
	}
	for _, url := range urls {
		go checkUrl(url)
	}
}

