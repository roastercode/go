package main

import (
	"fmt"
	"math"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	j := 2
	for j <= 20 {
		fmt.Print("\t", j)
		j = j + 5
	}
	fmt.Println()

	k := 3
	cnt := 30
	for k <= cnt {
		fmt.Println(k)
		k = k + 3
	}
	l := 1
	tnc := 12
	for l <= tnc {
		fmt.Print("\t", l)
		l = l + l
	}
	fmt.Println()

	m := 1
	nct := 12
	for m <= nct {
		fmt.Println(m)
		m++
	}
	hm := 12
	for n := 1; n <= hm; n++ {
		fmt.Print("\t", n)
	}
	fmt.Println()

	abc := 12
	for o := 5; o < abc; o++ {
		fmt.Println(o + 1)
	}
	for p := range "testing" {
		fmt.Print("\t", p)
	}
	println()

	var q float64 = 5.51
	var r int = 10

	println(int(q) + (r))
	// use math.Floor to return greatest integer value
	println(int(math.Floor(q+0.5)) + (r))

	println()

	// Print number of byte by character
	s := "Hello World"
	println(s[2])
	t := " "
	println(t[0])
}
