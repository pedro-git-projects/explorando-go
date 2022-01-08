package main

import "fmt"

// notice we don't need to declare each individual value as booelan
// noting a single bool will let the compiler infer the other value
// in () we put our return types, in this case, 3 boolean values
func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // xor
	comprarSorvete := trab1 || trab2
	
	return comprarTv50, comprarTv32, comprarSorvete

}

func main() {
	tv50, tv32, sorvete := compras(true, true)	
	fmt.Printf("Tv50: %t, Tv32: %t, Saudável: %t\n", tv50, tv32, !sorvete)
}
