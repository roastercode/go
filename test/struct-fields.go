/*

Struc Fields

Struct fields are accessed using a dot.

*/

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	v.Y = 6                // create the pointer to Y
	fmt.Println(v.Y * v.X) // use it with the X pointer
}
