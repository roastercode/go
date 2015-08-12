/*

Interfaces - interfaces.go

An interface type is defined by a set of methods.

A value of interface type can hold any value that implements those methods.

Note: There is an error in the example code on line 43. 'Vertex' (the value
type) doesn't satisfy 'Abser' because the 'Abs' method is defined only on
'*Vertex' (the pointer type).

First line was a = v and get the things wrong.
the a = &v version bring the things right

More information on interfaces and types :
https://golang.org/doc/effective_go.html#interfaces_and_types

*/

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not a *Vertex)
	// and does NOT implement Abser

	a = &v
	
	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
