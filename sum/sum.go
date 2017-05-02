package main

import "fmt"

func sum(bb []float64) float64 {
	total := 0.0
	for _, v := range bb {
		total += v
	}
	return total
}

// using sum function
func main() {
	nn := sum([]float64{98, 65 ,75 ,12 ,45})
	fmt.Println(nn)
}
