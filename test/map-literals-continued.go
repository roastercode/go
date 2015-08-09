/*

Map literals continued - map-literals-continued.go

If the top-level type is just a type name, you can omit it from the
elements of the literal.

*/

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

var n = map[string]Vertex{
	"Hackers Labs":   {41.3714, 8.5823},
	"Seti Institute": {37.3862296, -122.0513168},
}

func main() {
	fmt.Println(m)
	fmt.Println(n)
}
