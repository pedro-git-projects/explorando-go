package main

import "fmt"

func main() {
	var (
		a int
		b float64
		c bool
		d string
		e *int
	)

	// zero/garbage values in go 
	fmt.Printf("%d, %f, %t, %s, %p\n",a , b, c, d, e)

}
