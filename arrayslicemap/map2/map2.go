package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64 {
		"Pedro": 1700.99,
		"Raphael": 4999.99,
		"Rafael": 1200.00,
	}	

	funcsESalarios["Pedro"] = 1650.99
	fmt.Println(funcsESalarios)

	delete(funcsESalarios, "Não dá problema")

	  // chave valor
	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
