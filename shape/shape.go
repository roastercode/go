package main

// Structure and interfaces

import (
	"fmt"
	"math"
)

// INPUT DECLARATION
// SHAPE DECLARATION
type Shape interface {
	area() float64
	perimeter() float64
}

// CIRCLE STRUCT
type Circle struct {
	x, y, r float64
}

// CIRCLE AREA
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// CIRCLE PERIMETER
func (c *Circle) perimeter() float64 {
	return math.Pi * 2 * c.r
}


// RECTANGLE STRUCT
type Rectangle struct {
	x1, y1, x2, y2 float64
}

// RECTANGLE AREA
func (r *Rectangle) area() float64 {
	l := distance (r.x1, r.y1, r.x1, r.y2)
	w := distance (r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// RECTANGLE PERIMETER
func (r *Rectangle) perimeter() float64 {
	l := distance (r.x1, r.y1, r.x1, r.y2)
	w := distance (r.x1, r.y1, r.x2, r.y1)
	return l*2 + w*2
}

// DISTANCE DECLARATION
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}


// OUTPUT DECLARATION
func main() {
	var s Shape
	s = &Circle {0, 0, 5}         // using shape
	fmt.Println(s.perimeter())
	s = &Rectangle {0, 0, 10, 10} // using shape
	fmt.Println(s.perimeter())
}
