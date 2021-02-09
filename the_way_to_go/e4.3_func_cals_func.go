package main

var a string

func main() {
	a = "M"
	print(a)
	f()
}

func f() {
	a := "A"
	print(a)
	g()
}

func g() {
	print(a)
}
