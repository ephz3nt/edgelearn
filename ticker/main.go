package main

import (
	"fmt"
	"time"
)

func backgroundTask() {

	// this creates a new ticker which will
	// `tick` every 1 second
	ticker := time.NewTicker(1 * time.Second)

	// for every `tick` that our `ticker`
	// emits, we print `tock`
	for _ = range ticker.C {
		fmt.Println("tock")
	}
}

func main() {
	fmt.Println("Go Tickers Tutorial")

	go backgroundTask()
	// This print statement will be executed before
	// the first `tock` prints in the console
	fmt.Println("The rest of my application can continue")
	// here we use a empty select{} in order to keep
	// our main function alive indefinitely as it wuld
	// complete before our backgroundTask has a chance
	// to execute if we didn't
	select {}
}
