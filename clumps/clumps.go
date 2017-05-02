package main

import (
	"fmt"
)

func countClumps(xs []int) int {
	count := 0
	for i := 2; i < len(xs); i++ {
		if xs[i] != xs[i-1] && xs[i-1] == xs[i-2] {
			count++
		}
	}
	return count
}

func main() {
	clumps := []int{1, 1, 2, 3, 4, 4, 4, 4, 5, 6, 7, 7, 8}
	result := countClumps(clumps)
	fmt.Println(result)
}
