/*
Golang     : Slice v2
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
	var numbers []int

	printSlice(numbers)

	if(numbers == nil){
		fmt.Printf("slice is nil")
	}
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

