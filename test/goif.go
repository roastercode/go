package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 4 { // using 1 rather than 4 solve the stackoverflow
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
