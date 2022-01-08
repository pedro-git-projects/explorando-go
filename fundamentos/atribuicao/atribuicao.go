package main

import "fmt"

func main() {
	// keyword identifier type operator value
	var b byte = 3 //uint8
	fmt.Println(b)

	// identifier operator value
	i := 3
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i/= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2
	
	fmt.Printf("%d\n", i)

	// compount attribution
	x, y := 1, 2
	fmt.Printf("%d %d\n", x, y)
	
	// interesting value switch 
	x, y = y, x
	fmt.Printf("%d %d\n", x, y)

}
