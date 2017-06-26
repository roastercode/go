/*
Golang - Pointers v1
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
	var a int= 20              // actual variable declaration
	var ip *int                // pointer variable declaration

	ip = &a                    // store adress of a in a pointer variable

	fmt.Printf("Adress of a variable: %x\n", &a )

	// Adress stored in pointer variable
	fmt.Printf("Address stored in ip variable: %x\n", ip )

	// Access the value using the pointer
	fmt.Printf("Value of *ip variable: %d\n", *ip )
}

