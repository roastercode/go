/*
Golang - Scope II
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
var g int

func main() {

	/* local variable declaration */
	var a, b int

	/* actual initialization */
	a = 10
	b = 20
	g = a + b

	fmt.Printf("value of a = %d, b = %d and g = %d\n", a, b, g)
}

