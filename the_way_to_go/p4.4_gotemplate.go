package main

import "fmt"

var v int = 3

const c = "C"

type T struct {
	//...
}

func init() {
	//...
}

func main() {
	var a int
	Fun1()
	//...
	fmt.Println(a)
}

func (t T) Method1() {
	//...
}

func Fun1() {
	//...
}
