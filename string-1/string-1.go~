/*
Golang - Scope III
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

/* global variable declaration */
var a int = 20;

func main() {
	/* local variable declaration in main function */
	var a int = 10
	var b int = 20
	var c int = 0

	fmt.Printf("value of a in main() = %d\n", a);
	c = sum( a, b);
	fmt.Printf("value of b in main() = %d\n", c);
}

/* function to add two integers */
func sum(a, b int) int {
	fmt.Printf("value of a in sum() = %d\n", a);
	fmt.Printf("value of b in sum() = %d\n", b);

	return a + b;
}
