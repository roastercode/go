/*
Golang     : Structure v1
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
	var numbers = make([]int,3,5)

	printSlice(numbers)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

