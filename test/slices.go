/*

Slices - slices.go

A slice point to an array of values and also includes a length

'[]T' is a slice with elements of type 'T'.

*/

// Change from the original with double slice

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	e := []int{1, 4, 6, 8, 9, 10, 12}
	fmt.Println("s ==", s)
	fmt.Println("e ==", e)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("e(%d] == %d\n", i, e[i])
	}
}
