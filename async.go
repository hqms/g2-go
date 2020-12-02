package main

import (
	"fmt"
	"net/http"
)

type urlStatus struct {
	url    string
	status bool
}

func checkUrl(url string, c chan urlStatus) {
	_, err := http.Get(url)
	if err != nil {
		// The website is down
		c <- urlStatus{url, false}
	} else {
		// The website is up
		c <- urlStatus{url, true}
	}
}

func main() {

	urls := []string{
		"https://detik.com/",
		"https://xl.co.id/",
	}

	c := make(chan urlStatus)
	for _, url := range urls {
		go checkUrl(url, c)

	}
	result := make([]urlStatus, len(urls))
	for i, _ := range result {
		result[i] = <-c
		if result[i].status {
			fmt.Println(result[i].url, "is up.")
		} else {
			fmt.Println(result[i].url, "is down !!")
		}
	}

}
