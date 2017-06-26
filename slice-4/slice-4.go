/*
Golang     : Slice v4
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

	// Append allows nil slice
	numbers = append(numbers, 0)
	printSlice(numbers)

	// Add one element to slice
	numbers = append(numbers, 1)
	printSlice(numbers)

	// Add more than one element at a time
	numbers = append(numbers, 2,3,4)
	printSlice(numbers)

	// Create a slice numbers1 with double the capacity of earlier slice
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	// Copy content of numbers to numbers1
	copy(numbers1,numbers)
	printSlice(numbers1)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

