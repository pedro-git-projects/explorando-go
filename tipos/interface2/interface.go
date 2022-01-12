package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo string
	turbLigado bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turbLigado = true
}

func main() {
	// diretamente com o tipo podemos alterar diretamente o atributo
	carro1 := ferrari{"f40", false, 0}
	carro1.ligarTurbo()

	// ao trabalharmos com a interface precisamos passar a referência
	var carro2 esportivo = &ferrari{"f40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
	
}
