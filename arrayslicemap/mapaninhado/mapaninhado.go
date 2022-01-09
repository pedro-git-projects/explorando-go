package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64 {
		"P": {
			"Pedro Martins": 1700.00,
			"Pâmela Rocha": 3000.99,
		},
		"Q": {
			"Quênia Lima": 1234.98,
		},
		"R": {
			"Raquel Bragança": 5632.04,
		},
	}

	fmt.Println(funcsPorLetra)
	delete(funcsPorLetra, "P")
	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
