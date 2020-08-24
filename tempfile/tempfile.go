package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := ioutil.TempFile("car-images", "car-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(file.Name())
	fmt.Println(file.Name())
	if _, err := file.Write([]byte("Hello world\n")); err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
