package main

import "fmt"

// Recursion function
// Using uint64 because of big number
func factorial(x uint64) uint64 {
	if x == 0 {
		return 1
	}
	// 10 * 9 * 8 * 7 * 6 ... * 1
	return x * factorial(x-1)
}

/*
// Fibonnacci function
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
*/

// Another Fibonnacci function
func fib(n int) int{
	if n < 2{
		return n
	}
return fib(n-1) + fib(n-2)
}


func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))
	fmt.Println(factorial(20))
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fib(8))
	fmt.Println(fib(9))
	fmt.Println(fib(10))
}
