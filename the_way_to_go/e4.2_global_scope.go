package main

var a = "A"

func main() {
	f()
	g()
	f()
}

func f() {
	print(a)
}

func g() {
	a = "G"
	print(a)
}
