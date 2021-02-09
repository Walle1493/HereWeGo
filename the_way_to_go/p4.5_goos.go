package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Print(goos)
	path := os.Getenv("PATH")
	fmt.Println(path)
}
