package main

import (
	"fmt"
)

func notaParaMencao (n float64) string {
	if n >= 9 && n <= 10 {
		return "SS"
	} else if n >= 8 && n < 9 {
		return "MS"
	} else if n >=5 && n < 8 {
		return "MM"
	} else if n >= 3 && n <= 5 {
		return "MI"
	} else {
		return "II"
	} 
}

func main ()  {
	nota := 8.7
	fmt.Printf("Sua menção foi: %s\n", notaParaMencao(nota))
}
