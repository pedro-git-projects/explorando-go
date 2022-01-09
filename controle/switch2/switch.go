package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// A switch without a value will execute the first true case
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")

	}
}
