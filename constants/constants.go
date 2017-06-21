/*
    Golang - Experimentations
License    : GNU GPL v3 or later
Author     : Aurélien DESBRIERES
Mail       : aurelien@hackers.camp
Project    : Python Experimentations
Created on : June 2017

Write with Emacs-Nox ──────────────┐
Go Constants ──────────────────────┘
*/

package main

import "fmt"

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("value of area : %d\n", area)
}
