package main

// creating an import alias
import (
	"fmt"
	alias "math"
)

func main() {
	// type qualifier identifier type value 
	const PI float64 = 3.1415

	var raio = 3.2 // type (float64) inferred by the compiler 

	// go style of creating variables
	area := PI *alias.Pow(raio, 2)
	fmt.Printf("A área da circunferência é %.2f\n", area)

	// defining variables inside blocks
	const (
		a = 1
		b = 2
	)

	var (
		c = 3 
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println("Declaring many variables in the same line, for instance: ", e, f)

	g, h, i := 2, false, "cool"
	fmt.Println("Declaring variables of different tpes using the := operator: ", g, h, i)
}
