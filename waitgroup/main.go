package main

import (
	"fmt"
	"sync"
)

func myFunc(wg *sync.WaitGroup) {
	fmt.Println("Inside my goroutine")
	wg.Done()
}

func main() {

	fmt.Println("Hello")
	var wg sync.WaitGroup
	wg.Add(2)
	go myFunc(&wg)
	go func() {
		fmt.Println("Inside my goroutine2")
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("Finished Execution")
}
