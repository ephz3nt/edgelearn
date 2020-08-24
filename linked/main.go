package main

import (
	"container/list"
	"fmt"
)

func main() {
	myList := list.New()
	myList.PushBack(1)
	myList.PushFront(2)
	for element := myList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
