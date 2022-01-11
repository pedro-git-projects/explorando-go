package main

import "fmt"

// creating a zero parameter function with return type of function
func colusure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	// will print the main scope
	fmt.Println(x)

	impremeX := colusure()
	// will print lexical scope
	impremeX()
}
