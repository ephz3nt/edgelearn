package main

import "fmt"

func getLimit() func() int {
	limit := 10
	return func() int {
		limit -= 1
		return limit
	}
}

func main() {
	limit := getLimit()
	fmt.Println(limit())
	fmt.Println(limit())
	limit2 := getLimit()
	fmt.Println(limit2()) // 9
	fmt.Println(limit2()) // 8
	fmt.Println(limit())
}
