package main

import (
	"fmt"
	"time"
)

// Go tem apenas o for como laço de repetição
func main() {
		
	// While-like for
	// while i <= 10
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	// mixing things up
	for i := 1; i <= 10; i++ {
		if i %2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Ímpar ")
		}
	}

	// infinite loop
	for {
		fmt.Println("Infinito")
		time.Sleep(time.Second * 3)
	}
}
