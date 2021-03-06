// leaving out the length 4 in [4] uses  slice instead of an array
// which is generally better for performance
package main

import "fmt"

func mapFunc(fn func(int) int, slice []int) {
	for i := range slice {
		slice[i] = fn(slice[i])
	}
	//return slice
}

func mul(n int) (m int) {
	m = n * 10
	return
}

func main() {
	slice := []int{1, 2, 3, 5, 4, 6, 8, 7}
	mapFunc(mul, slice)
	fmt.Println(slice)
}
