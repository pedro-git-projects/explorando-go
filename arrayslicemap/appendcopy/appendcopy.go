package main

import "fmt"

func main() {
	array1 := [3]int{1,2,3}
	var slice1 []int

	//can't append to arrays array = append(array1,4,5,6)
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	// [recebe] [atribui]
	copy(slice2, slice1)
	// O copy não faz com que o slice cresça 
	fmt.Println(slice2)
}
