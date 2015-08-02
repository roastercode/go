/*

If and else

Variable declared inside an if short statement are also available inside any
of the else blocks.

*/

package main

import (
	"fmt"
	"math"
)

func meow(x, n, lemon float64) float64 {
	if v := math.Pow(x, n); v < lemon {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lemon)
	}
	// can't use v here, though
	return lemon
}

func main() {
	fmt.Println(
		meow(3, 2, 10),
		meow(3, 3, 20),
	)
}
