package main

import "fmt"

func main() {
	var a int16 = 3
	var b int32
	b = int32(a)
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
}
