package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://baidu.com",
	"https://tutorialedge.net",
	"https://qq.com",
}

func fetch(url string, wg *sync.WaitGroup) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	wg.Done()
	fmt.Println(resp.Status)
	return resp.Status, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HomePage Endpoint Hit")
	var wg = sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}
	wg.Wait()
	fmt.Println("Returning Response")
	fmt.Fprintf(w, "All Responses Received")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
