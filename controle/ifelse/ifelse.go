package main

import "fmt"

func imprimirResultado(nota float64) {
	// A expressão não tem parênteses
	// expressões de uma linha também devem ter parênteses
	if nota >=7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(7.3)
}
