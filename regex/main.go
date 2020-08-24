package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	input := "/usr/bin/kubectl"

	resolvedPath, _ := filepath.EvalSymlinks(input)

	fmt.Printf("Input: %v :%v\n", input, resolvedPath)
}
