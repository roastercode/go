/*

Exercice: Fibonacci closure

Let's have some fun with functions

Implement a fibonacci function that returns a function (a closure)
that returns successive fibonnacci numbers.

*/

package main

import "fmt"

// fibonnacci is a function that returns
// a function that returns an int.

func fibonnacci() func() int {
	x := 0
	y := 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	f := fibonnacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
