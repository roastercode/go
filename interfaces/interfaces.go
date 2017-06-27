/*
Golang     : Intefaces v1
License    : GNU GPL v3 or later
Author     : Aurélien DESBRIERES
Mail       : aurelien@hackers.camp
Project    : Golang Experimentations
Created on : June 2017

Write with Emacs-Nox ──────────────┐
Go lang ───────────────────────────┘
*/

package main

import (
	"fmt"
	"math"
)

// Define an interface
type Shape interface {
	area() float64
}

// Define a circle
type Circle struct {
	x,y,radius float64
}

// Define a rectangle
type Rectangle struct {
	width, height float64
}

// Define a method for circle (implementation of Shape.area())
func(circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

// Define a method for rectangle (implementation of Shape.area())
func(rect Rectangle) area() float64 {
	return rect.width * rect.height
}

// Define a method for shape
func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{x:0,y:0,radius:5}
	rectangle := Rectangle {width:10, height:5}

	fmt.Printf("Circle area: %f\n",getArea(circle))
	fmt.Printf("Rectangle area: %f\n",getArea(rectangle))
}

