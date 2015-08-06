/*

Slicing slices - slicing-slice.go

Slices can be re-sliced, creating a new slice value that points
to the same array.

The expression

 's[lo:hi]

evaluates to a slice of the elements from 'lo' through 'hi-1', inclusive.
Thus

 's[lo:lo]

is empty and

 s[lo:lo11]

has one element.

*/

// Change from the original with a double slice using defer
// created by that way in inverted result

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	e := []int{1, 4, 6, 8, 9, 10}
	fmt.Println("s ==", s)
	defer fmt.Println("e ==", e)
	fmt.Println("s[1:4] ==", s[1:4])
	defer fmt.Println("e[1:4] ==", e[1:4])

	// missing low index implies 0
	fmt.Println("s[:3] ==", s[:3])
	defer fmt.Println("e[:3] ==", e[:3])

	// missing high index implies len(s)
	fmt.Println("s[4:] ==", s[4:])
	defer fmt.Println("e[4:] ==", e[04:])
}
