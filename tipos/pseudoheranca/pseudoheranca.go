package main

import "fmt"

// Go não possui herança, mas somos capazes de simular herança através de composição utilizando
// campos anônimos

type carro struct {
	nome string
	velocidadeAtual int
}

type ferrari struct {
	carro // campo anônimo
	turboLigado bool
}

func main() {
	f := ferrari{}
	// note que ferrari não possui campos nome e velocidadeAtual
	// mas por conta do campo anônimo eles se tornam acessíveis
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com o turbo ligado? %t\n", f.nome, f.turboLigado)
	
}
