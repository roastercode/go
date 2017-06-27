/*
Golang     : Recursion v0
License    : GNU GPL v3 or later
Author     : Aurélien DESBRIERES
Mail       : aurelien@hackers.camp
Project    : Golang Experimentations
Created on : June 2017

Write with Emacs-Nox ──────────────┐
Go lang ───────────────────────────┘
*/

package main

import "fmt"

func factorial(i int)int {
	if(i <= 1) {
		return 1
	}
	return i * factorial(i - 1)
}

func main() {
	var i int = 15
	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))
}
