package main

import "fmt"

type (
	IZ  int
	FZ  float32
	STR string
)

func main() {
	var a = 3.0
	fmt.Printf("%T\n", a)
	b := int(a)
	fmt.Printf("%T\n", b)
	c := IZ(a)
	fmt.Printf("%T\n", c)
}
