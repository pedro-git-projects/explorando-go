package main

import "fmt"

func fatorial(n int) (int, error) {
	// enters first true case
	switch {
		case n < 0:
			return -1, fmt.Errorf("número inválido: %d", n)
		case n == 0:
			return 1, nil
		default:
			fatorialAnterior, _ := fatorial(n -1)
			return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-2)
	if err != nil {
		fmt.Println(err)
	}
}
