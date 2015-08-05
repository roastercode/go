// Struct literal X3 Vertex struct
package main

import "fmt"

type Vertex struct {
	X, Y, Z int
}

var (
	v1 = Vertex{1, 2, 3}    // has type Vertex
	v2 = Vertex{X: 1, Z: 3} // Y:0 is implicit
	v3 = Vertex{}           // X:0, Y:0 and Z:0
	p  = &Vertex{1, 2, 3}   // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
