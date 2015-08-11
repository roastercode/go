/*

Methods - methods.go

Go does not have classes. However, you can define methods on stuct types.

The method receiver appears in its own argument list between the func keyword and the method name.

*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y, Z float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func main() {
	v := &Vertex{3, 4, 5}
	fmt.Println(v.Abs())
}
