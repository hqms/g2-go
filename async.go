package main

import (
	"fmt"
	"net/http"
	"sync"
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
	var wg sync.WaitGroup

	for _, u := range urls {

		wg.Add(1)
		go func(url string) {

			defer wg.Done()

			checkUrl(url)
		}(u)
	}

	wg.Wait()
}
