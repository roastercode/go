/*

Adding elements to a slice

It is common to append new elements to a slice, and so Go provides
a built-in 'append' function. The documentation of the built-in package
describes 'append'.

 'func append(s []T, vs ...T) []T

The first parameter s of append is a slice of type T, and the rest are T
values to append to the slice.

The resulting value of append is a slice containing all the elements of
the original slice plus the provided values.

If the backing array of 's' is too small to fit all the given values a bigger
array will be allocated. The returned slice will point to the newly allocated
array.

(To learn more about slices, read the Slice: usage and internals article.)

https://botbot.me/freenode/go-nuts/

*/

package main

import "fmt"

func main() {
	var a []int
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a)

	// use for statement i := 5; i < 20; i++ {} to specifies repeated execution of the block
	for i := 5; i < 20; i++ {
		a = append(a, i)
		printSlice("a", a)
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
