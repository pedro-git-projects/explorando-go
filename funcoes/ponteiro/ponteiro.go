package main

import "fmt"

// go makes use of pure functions
// this means that functions are side effect free
func inc1(n int) {
	n++
}

// to add side effects one must use pointers
func inc2 (n *int) {
	//dereferencing and incrementing
	*n++ 
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)
	
	inc2(&n)
	fmt.Println(n)

}
