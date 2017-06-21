/*
Golang - Scope I
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

func main() {
	/* local variable declaration */
	var a, b, c int

	/* actual initialization */
	a = 10
	b = 20
	c = a + b

	fmt.Printf ("value of a = %d, b = %d and c = %d\n", a, b, c)
}

