package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 1 //int64	
	
	var p *int = nil // declarando um ponteiro e inicializndo como nulo

	p = &i // atribuindo  endereço da variável i com o operador &
	*p++ // derreferenciando p e incrementando seu valor
	i++ // incrementando diretamente pela variável
	
	fmt.Printf("i: %d, *p: %d, &i: %p, p: %p\n", i, *p, &i, p)
	
	aString := "string"
	shortPointerDeclaration := &aString
	fmt.Println("Declaring and initializing a ", reflect.TypeOf(shortPointerDeclaration))

	// Go NÃO tem aritmética de ponteiros!
	//	p++
}
