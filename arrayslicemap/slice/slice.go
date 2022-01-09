package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3] int{1,2,3} //array
	s1 := [] int{1,2,3} // slice

	// slices are references to arrays
	fmt.Printf("a1 is an array and has type of: %s and value: %d \ns1 is a slice has type of: %s and value: %d\n", reflect.TypeOf(a1), a1, reflect.TypeOf(s1), s1)

	a2 := [5] int{1,2,3,4,5} // array!
	// Um slice não é um array, um slice define um pedaço de array
	s2 := a2[1:3] // do índice 1(inclusivo) até o índice 3 (exclusivo)
	// Um slice não cria um novo array, ele aponta para uma parte específica do array
	fmt.Println(a2, s2)

	s3 := a2[:2] // do início
	fmt.Println(s3)
	// um slice tem um tamanho e um ponteiro que pega de um endereço até o fim do array através de seu tamanho
}
