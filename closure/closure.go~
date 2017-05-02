package main

import "fmt"

// first version
func half(n int) (int, bool) {
	if n % 2 == 0 {
		return n/2, true
	} else {
		return n/2, false
	}
}

// simpler version
func halfer(n int) (int, bool) {
	return n/2, n%2 == 0
}


// both function are same giving back: <x> false
// or <x> true
/*
func main() {
	h, even := half(4)
	fmt.Println(h, even)
	g, even := halfer(4)
	fmt.Println(g, even)
}
*/

// Using separate functions
func calculus() {
	h, even := half(4)
	fmt.Println(h, even)
}

func calcula() {
	g, even := halfer(3)
	{
		g := 5
		fmt.Println(g)
	}
	fmt.Println(g, even)
}

func main(){
	calculus()
	calcula()
}

