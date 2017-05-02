package main

import "fmt"

func main() {
	// Different type of declaration
	var x, y = 5, 2
	fmt.Println(x % y)
	
	// Another way of declaration
	w, z := 5, 2
	fmt.Println(w % z)

	// Yet another type of declaration and way to print the result
	var str = "Hello World " + "test"
	fmt.Println(str)
	fmt.Println(len(str))
}
