package main

import (
	"fmt"
)

func main() {

	// MAPS IS A DICTIONNARY

	x := make(map[string]int)
	x["dico"] = 10
	x["something"] = 0

	fmt.Println(x["dico"])
	fmt.Println(x)
	fmt.Println(x["not match"]) // because it does not match "dico" it print 0

	v, ok := x["something else"]
	fmt.Println(v, ok) // print 0 false because x["something"] = 0 == does not match
	z, ok := x["something"]
	fmt.Println(z, ok) // print 0 true because x["something"] = 0 == match

	fmt.Println("*\tPRINTING INT IN MAP\t*\n")
	y := make(map[int]int)
	y[1] = 10
	y[2] = 0

	w, ok := y[1]
	fmt.Println(w, ok)
	b, ok := y[2]
	fmt.Println(b, ok)
	fmt.Println(y)

	fmt.Println("*\tCLEANER ARRAY SOLUTION\t*")

	a:= map[int]int{
		7: 10,
		3: 0,
	}
	fmt.Println(a)

	fmt.Println("*\tMOVE THINGS USING DELETE\t*")

	delete(a, 7)
	fmt.Println(a)
}
