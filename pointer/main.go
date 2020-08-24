package main

import "fmt"

func main() {
	var age *int
	age = new(int)
	*age = 26

	fmt.Println(*age)
	fmt.Println(&age)
}
