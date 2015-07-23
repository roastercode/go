/*
Basic types

Go's basic types are

bool - boolean

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

*/

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	a    bool
/* http://golang.org/ref/spec#The_zero_value
in The Tour of Go, they specify '= false'
The point is that is for a better reading, but,
the zero value is 'false' by default.
The use of 'bool = false' is for the reader, a standart of writing.  */
	MaxInt  uint64      = 1<<64 - 1
	z       complex128  = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, a, a)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
