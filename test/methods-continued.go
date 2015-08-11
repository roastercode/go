/*

Methods continued - methods-continued.go

You can declare a method on any type that is declared in your package,
not just struct types.

However, you cannot define a method on a type from another package
(including built in types)

*/

package main

import (
	"fmt"
	"math"
)

type MyFloat float64
type MyInt int

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (i MyInt) Abs() int {
	if i < 0 {
		return int(-i)
	}
	return int(i)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	i := MyInt(6 * f)
	fmt.Println(f.Abs())
	fmt.Println(i.Abs())
}
