package main

import (
	"fmt"
)

func main() {
	// 1st try
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	// 2nd try
	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83

	var total float64 = 0

	// Two way to make the same declaration that will double the result
	for _, a := range y {
		total += a
	}
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println(total / float64(len(y)))

	//3rd try (clearner version) with renamed variables
	//Change the division by a multiplication
	//Using direct delclaration of := rather than var
	z := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	var maximus float64 = 0

	// Two way to make the same declaration that will double the result
	for _, b := range z {
		maximus += b
	}
	for c := 0; c < len(z); c++ {
		maximus += z[c]
	}
	fmt.Println(maximus * float64(len(z)))


	// Multidimentional arrays
	n := [6][4]float64{       // First  [array] declared the number of time
		
		{ 1, 2, 3, 4 },
		{ 2, 3, 4, 5 },
		{ 3, 4, 5, 6 },
		{ 7, 8, 9, 10 },
	}
	m := [4][6]float64{
		
		{ 1, 2, 3, 4 },
		{ 2, 3, 4, 5 },
		{ 3, 4, 5, 6 },
		{ 7, 8, 9, 10 },
	}

	o := [4][4]float64{       // Second [array] declared the number of delclaration in the array
		
		{ 1, 2, 3, 4 },
		{ 2, 3, 4, 5 },
		{ 3, 4, 5, 6 },
		{ 7, 8, 9, 10 },
	}
	
	p := [4][4]float64{
		
		{ 1, 2, 3, 4 },
		{ 2, 3, 4, 5 },
		{ 3, 4, 5, 6 },
		{ 7, 8, 9, 10 },
	}

	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(o)
	fmt.Println(p[0][0])     // Request to print for a specific array
	p[0][1] = 22             // Change the value of an array (here the second number in the first arry
	fmt.Println(p)
	o[1][1] = 44             // Changing the value of the second arry second number
	fmt.Println(o)


	// 4th try - Creating a slice with an empty array
	q := []float64{1, 2, 3, 4}

	q = make([]float64, 10, 15) // The second number exclude the 5 more number
	
	r := q[1:]
	fmt.Println(q, r)

	// 5th try - Another Slice
	
	xs := []int{1,2,3}
	fmt.Println(xs)

	xs = append(xs, 4) // Print one more (append put in the next one)
	fmt.Println(xs)

	xs = append(xs, 5) // Print one more on the one added before
	fmt.Println(xs)

	xs = append(xs, 10) // Print +10
	fmt.Println(xs)

	xs = append(xs, 4, 5, 6, 7) // Print four more
	fmt.Println(xs)

	sx := []int{1,2,3}
	fmt.Println(sx)
	sx = append(sx, 10) // Print + the 10th
	fmt.Println(sx)

	fmt.Println("--------------------------")

	ys := make([]int, 4)
	copy(ys[2:], xs)        // copy everything inside the [xs and ys] arrays
	fmt.Println(ys)

	yx := make([]int, 4)
	copy(yx, xs)        // copy everything inside the [xs and ys] arrays
	fmt.Println(yx)

	yz := make([]int, 4)
	yz[0] = 8

	xa := make([]int, 4)
	xa[0] = 8
	
	xs = append(xs, yz[0]) // include the yz array as it
	fmt.Println(xs)

	xs = append(xs, yz...) // include the yz array at the end of the xz array the '.' will give 0 0 0
	fmt.Println(xs)
	
	xa = append(sx, xa[0]) // the [0] will request for the yz array only
	fmt.Println(sx)

	fmt.Println("\n\n*\tLAST BUT NOT LEAST\t*\tAND CLEAN\t*\n")
	
	xw := []int{1, 2, 3}
	fmt.Println("*", xw, "*")
	xw = append(xw, 4, 5, 6, 7)
	fmt.Println("*", xw, "*")

	yb := make([]int, 4)
	yb[0] = 8
	yb[1] = 9

	xw = append(xw, yb[0:3]...) // + the 2nd number in the array to complet to 10 here 3
	fmt.Println("*", xw, "*")
}
