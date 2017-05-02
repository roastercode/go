package main

import (
	"fmt"
	"sort"
)

func main() {

	// INTRODUCTION
	fmt.Println("*\tMAP FINDER\t*")

	// Map of int
	xs := []int{
		48,96,86,68,
		57,82,63,78,
		37,34,83,27,
		19,97, 9,17,
	}

	min := xs[0]
	for _, v := range xs {
		if v < min {
			min = v
		}
	}
	fmt.Println(min)

	// Map of float
	xy := []float32{
		5.5,8.5,5.7,
		1.2,2.3,1.5,
		1.7,3.1,3.2,
		2.4,6.5,6.4,
	}

	max := xy[0]
	for _, w := range xy {
		if w > max {
			max = w
		}
	}
	fmt.Println(max)
	
	// ADDING SCANF FUNCTION
	fmt.Printf("Give me the number from 0 to 16\nin the map you are looking for: ")
	// ADDING THE SYMBOL STRING FOR THE SCANF FUNCTION
	var number int

	// CALL OF THE SYMBOL
	fmt.Scanf("%d", &number)

	fmt.Println("This is what you are looking for in the map:", xs[number])
	// RETURN THE RIGHT ANSWER


	fmt.Println("*\t LET'S TEST WITH SORT PACKAGE NOW!\t*")

	sort.Ints(xs)

	// Printing two map in a time
	fmt.Println(xs[1], xs[2])

	// Map of string
	var code string
	fmt.Printf("Give me the code state you are looking for: ")
	fmt.Scanln(&code)

	states := map[string]string{
		"CA": "California",
		"NE": "Nevada",
	}
	fmt.Println(states[code])
}
