package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var str = "Hello World \n	test"
	fmt.Println(len(str))
	fmt.Println(str)

	var string = "Hello World " + "test"
	fmt.Println(len(string))
	fmt.Println(string)

	var x int = 5
	var y uint64 = 6754453245675543
	var z float64 = 6754453245675543.7844514
	var strong = "Hello World " + strconv.Itoa(x) + fmt.Sprint(y)
	fmt.Println(strong)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(uint64(z) / (y))
	fmt.Println(uint64(z) + (y))
	fmt.Println(float64(x) + (z))

	var v float64 = 5.9
	var w int = 10
	fmt.Println(int(v) + (w))

	var a float64 = 5.9
	var b int = 10

	fmt.Println(int(math.Floor(a+0.5)) + (b))
}
