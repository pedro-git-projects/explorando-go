package main

import "fmt"

// file scope function variable
var soma = func (a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2,3))	

	// local function variable
	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2,3))
}
