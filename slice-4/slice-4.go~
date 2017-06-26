/*
Golang     : Slice v3
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
	// Create a slice
	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	// Print the original slice
	fmt.Println("numbers ==", numbers)

	// Print the sub slice starting from index 1(included) to index 4(excluded)
	fmt.Println("numbers[:3] ==", numbers[:3])

	// Missing upper bound implies len(s)
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	// Print the sub slice starting from index 0(included) to index 2(excluded)
	number2 := numbers[:2]
	printSlice(number2)

	// Print number3 := numbers[2:5]
	number3 := numbers[2:5]
	printSlice(number3)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

