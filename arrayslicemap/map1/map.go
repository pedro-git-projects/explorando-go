package main

import "fmt"

func main() {
	//keyword identifier keyword[chave]valor
	// var aprovados map[int]string	

	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[12345678909] = "Ana"
	aprovados[12345678912] = "Juca"

	fmt.Printf("%v", aprovados)
	
	for cpf, nome := range aprovados {
		fmt.Printf("%s, (CPF: %d)\n", nome, cpf)
	}

	// lendo através de uma chave
	fmt.Println(aprovados[12345678978])

	delete(aprovados, 12345678912)
	fmt.Println(aprovados,12345678912)
}
