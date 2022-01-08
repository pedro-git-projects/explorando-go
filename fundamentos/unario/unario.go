package main

import "fmt"

func main() {
	x := 1
	y := 2

	fmt.Printf("x: %d, y: %d\n", x, y)
	// apenas notação pósfixa!
	x++ // x += 1 ou x = x + 1
	y-- // y -=1

	fmt.Printf("x: %d, y: %d\n", x, y)

}
