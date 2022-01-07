package main

import "fmt"

func main() {
	fmt.Print("Same")
	fmt.Print(" line\n")

	fmt.Println("New")
	fmt.Println("line")

	x := 3.141516

	fmt.Printf("The value of x is %f\n", x)

	// transforming other types to strings
	xs := fmt.Sprint(x)

	fmt.Printf("The value of x is %s\n", xs)

	fmt.Print("The value of x is " + xs + "\n")

	fmt.Println("The value of x is", x)
	
	// format types in go
	a := 1
	b := 1.9999
	c := false
	d := "string"
	fmt.Printf("%d %f %t %s\n", a, b, c, d)
	// generic formating 
	fmt.Printf("%v %v %v %v\n", a, b, c, d)
}
