package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 // float64
	y := 2 // int

	fmt.Println(x / float64(y))

	// be careful with directly casting integers to strings in go 
	fmt.Println("A string " + string(123))

	// it is best to use:
	fmt.Println("Better casting " + strconv.Itoa(123))

	// string to int
	// this method returns a value and an erro
	// we're ignoring the error ATM
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// string to bool
	// all other values are parsed as false
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("True")
	}

	b2, _  := strconv.ParseBool("1")
	if b2 {
		fmt.Println("True")
	}
}
