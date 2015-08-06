/*

Making slices - making-slices.go

Slices are created with the 'make' function . It works by allocating
a zeroed array and retruning a slice that refers to that array:

 'a := make([]int, 5)		// len(a)=5'

To specify a capacity, pass a third argument to 'make':

 'b :=([]int, 0, 5)			// len(b)=0, cap(b)=5'

 'b = b[:cap(b)]				// len(b)=5, cap(b)=5'
 'b = b[1:]							// lenb(b)=4, cap(b)=4'

*/

package main

import "fmt"

func main() {
	a := make([]int, 6)
	printSlice("a", a)
	b := make([]int, 0, 6)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:6]
	printSlice("d", d)
	e := make([]int, 3, 5)
	printSlice("e", e)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// Added to note func len and cap
// http://golang.org/pkg/builtin/#len
// http://golang.org/pkg/builtin/#cap
