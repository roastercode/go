/*

Defer

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not 
executed until the surrounding function returns.

*/


package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	defer fmt.Printf("Hello World!\n\n")

var (
	a        bool
	MaxInt   uint64     = 1<<64 - 1
	z        complex128 = cmplx.Sqrt(-5 +12i)
)

	const f = "\n%T(%v)\n"
	fmt.Printf(f, a, a)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
