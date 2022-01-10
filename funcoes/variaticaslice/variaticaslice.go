package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	// a variatic function NEEDS to recieve a slice instead of an array
	aprovados := []string{"Pedro", "Tati", "Raphael"}	
	// ... will spread the slice values as parameters
	imprimirAprovados(aprovados...)
}
