package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	// interfaces vazias permitem que criemos tipos genéricos (mais ou menos como object no java)
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)
	
	type dinamico interface{}
	var coisa2 dinamico = "opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Explorando golang"}
	fmt.Println(coisa2)
}
