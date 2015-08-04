package main

import (
	"testing"
	"fmt"
	"gonum"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*3.14 + 1 }
func needFloat(x float64) float64 {
	return x * 6.28
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
