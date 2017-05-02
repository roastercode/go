package main

// Power // Exponent function

import (
	"fmt"
	"math"
)

type Function func(x, y float64) float64

func exponent(x,y float64 ) float64 {
	return math.Pow(x, y)
}

func takesFunction(f Function, x float64, y float64) float64{
	return f(x,y)
}

func main() {
	fmt.Println(exponent(3, 3))
}
