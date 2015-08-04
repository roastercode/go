/*

Pointers to structs

Struc field can be accessed through a struct pointer.
The indirection through the pointer is transparent.

*/

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Botox struct {
	Z int
	W int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	u := Botox{3, 4}
	o := &u
	o.Z = 2e3
	fmt.Println(p.X * o.Z)
}
