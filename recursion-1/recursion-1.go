/*
Golang     : Recursion v1
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

func fibonaci(i int) (ret int) {
	if i == 0 {
		return 0
	}

	if i == 1 {
		return 1
	}

	return fibonaci(i-1) + fibonaci(i-2)
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d \n", fibonaci(i))
	}
}

