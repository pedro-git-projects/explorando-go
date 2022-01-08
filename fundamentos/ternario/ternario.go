package main

import "fmt"

// There is no ternary operator in go

func obterResultado (nota float64) string {
	// return nota >= 6 ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}


func main() {
	fmt.Printf("%s\n", obterResultado(8.5))
}
