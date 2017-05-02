package main

import (
	"fmt"
	"os"
)
// different type of variables
var w string
var x string = "Hello World"
// empty variable
var y string
var z string = "Just as I want it!"
const milesToKilometers = 1.60934
const pTkg = 0.453592

func main() {
	// print a line
	fmt.Println(x)
	// print a blank line because y is an empty variable
	fmt.Println(y)
	// print a line
	fmt.Println(z)
	// change the value of z and print it in a line
	z = "Wow so good Go!"
	fmt.Println(z)
	// short end syntax to call a variable (only available inside a function)
	x := "1 point!"
	fmt.Println(x)
	// redefine an empty variable
	y := 2
	fmt.Print(y)
	w := " for the dog!"
	fmt.Println(w)
	{
		var y = 10
		fmt.Println(y)
	}
	{
		// a constant can't be change in same semicolon
		const x = 5
		fmt.Println(x)
	}
	// constant declaration
	const (
		a int = 5
		b int = 6
		c int = 15
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var name = "aurelien"
	fmt.Println("Hello my name is", name)

	var noom string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &noom)

	var d = 5
	var e = 10
	fmt.Println(d * e)

	var miles float64
	fmt.Print("Enter miles: ")
	fmt.Scanf("%f", &miles)
	kilometers := miles * milesToKilometers
	fmt.Println("Kilometers", kilometers)

	var fahrenheit float64
	fmt.Print("Enter number(fahrenheits): ")
	fmt.Scanln(&fahrenheit)
	celsius := (fahrenheit - 32) * (5.0 / 9)
	fmt.Println(celsius,"°C")

	var weight float64
	fmt.Print("Enter number(pounds): ")
	fmt.Scanln(&weight)
	lbs := weight * pTkg
	fmt.Println(lbs, "kg")
	
	os.Exit(1) // Exit for os import without error
}
