// array.go
/*

Arrays

The type [n]T is an arry of n values of type T.

The expression

var a [10]int
declares a variable a as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized.
This seems limiting, but don't worry; Go provides a convenient way of
working arrays

*/

package main

import "fmt"

func main() {
	var a [3]string
	a[0] = "Hello"
	a[1] = "World"
	a[2] = "!"
	fmt.Println(a[0], a[1], a[2])
	fmt.Println(a)
}
