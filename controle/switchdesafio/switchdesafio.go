package main

import "fmt"

func main() {
	nota := 8 

	switch {
	case nota >= 9 && nota <= 10:
		fmt.Println("A")
	case nota >= 8 && nota < 9:
		fmt.Println("B")
	case nota >=5 && nota < 8:
		fmt.Println("C")
	case nota >= 3 && nota < 5:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
