package main

import "runtime/debug"

func f3() {
	debug.PrintStack()
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	// A computer program is an execution from the stack
	// in this case, we call main, before main ends, it calls f1
	// before f1 resolves, it calls f2 and places it into the stack
	// f1 calls f2 and so on
	f1()	
}
