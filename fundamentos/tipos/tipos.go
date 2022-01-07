package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// integer lietrals
	integerLiteral := 666
	fmt.Println(1, 2, 1000)	
	fmt.Println("A literal integer has type", reflect.TypeOf(3200))
	fmt.Println("A literal integer has type", reflect.TypeOf(integerLiteral))

	// unsigend integers uint8 -> byte , uint16 -> short, uint32 -> int, uint64 -> long
	var b byte = 255 // byte is an alias to uint8
	fmt.Printf("A byte has type %s\n", reflect.TypeOf(b))

	// signed int8 int16 int32 int64 
	i1 := math.MaxInt64
	fmt.Println("The max value of int types is", i1)
	fmt.Println("i1 has type", reflect.TypeOf(i1))

	// rune is an alias for integer char values (unicode)
	var i2 rune = 'a'
	fmt.Println("A rune has type of", reflect.TypeOf(i2))
	fmt.Println("The value of i2 is:", i2)

	// real numbers (float32, float64)
	var x float32 = 49.99 
	fmt.Println("The type of x is, ", reflect.TypeOf(x))
	fmt.Println("The literal 49.99", reflect.TypeOf(49.99), " which is the default")

	// boolean 
	bo := true
	fmt.Println("bo has type", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Hello, my name is Pedro"
	fmt.Println(s1 + "!!!")
	fmt.Println("The s1 string size is ", len(s1))

	// string with multiple lines
	s2 := `Hello
	my
	name
	is 
	Pedro`

	fmt.Println("The s2 string size is ", len(s2))

	// There  is no char type in c, chars are int32
	char := 'a'
	fmt.Println("There is no char type in go, only", reflect.TypeOf(char))
}
