package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Printf("Datas, d1: %s, d2: %s\n", d1, d2)
	fmt.Printf("Even though this are date objects, comparing them will result in %t because we're comparing values and not adresses\n", d1 == d2) 
	fmt.Println("Using the equals method will have the same behaviour, resulting in", d1.Equal(d2))

	// creating a struct
	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Pedro"}
	p2 := Pessoa{"Pedro"}

	fmt.Printf("Comparing two different structs will compare their values, and not their references! So p1 and p2 results in %t\n", p1 == p2)

}
