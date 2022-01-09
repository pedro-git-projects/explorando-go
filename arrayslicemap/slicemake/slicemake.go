package main

import "fmt"

func main() {
	s := make([]int, 10) // creates a slice with 10 elements	
	s[9] = 12 // altering index 9 element to 12
	fmt.Println(s)

	s = make([]int,10, 20) // o slice terá dez elementos mas o ARRAY pra o qual slice aponta terá 20 posições
	fmt.Println(s, len(s), cap(s)) // capcity will return the capacity of the slices's internal array

	s = append(s,1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))
	
	// overflowing the array a slice points to will create a bigger array
	// instead of throwing an error
	s = append(s,1)
	fmt.Println(s, len(s), cap(s))
}
