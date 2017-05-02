package main

// Structure and interfaces

import (
	"fmt"
	"math"
)

// INPUT DECLARATION
// POINT DECLARATION
type Point struct {
	x, y float64
}

func (p Point) distance(p2 Point) float64 {
	a := p2.x - p.x
	b := p2.y - p.y
	return math.Sqrt(a*a + b*b)
}


// RECTANGLE DECLARATION
type Rectangle struct {
	p1, p2 Point
}

func (r *Rectangle) area() float64 {
	l := r.p1.distance(Point{r.p1.x, r.p2.y})
	w := r.p1.distance(Point{r.p2.x, r.p1.y})
	return l * w
}


// CIRCLE DECLARATION
type Circle struct {
	Point
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}


// CUBE DECLARATION
type Cube struct {
	n int
}

func (k *Cube) area() int {
	return k.n * k.n * k.n
}
	


// OUTPUT DECLARATION
func main() {

	// RECTANGLE OUTPUT
	r := &Rectangle {
		p1: Point {0, 0},
		p2: Point {10, 10},
	}
	fmt.Println(r.area())

	r = &Rectangle {
		p1: Point {1, 1},
		p2: Point {0, 0},
	}
	fmt.Println(r.area())

	
	// CIRCLE OUTPUT
	c := &Circle {
		Point: Point {0, 0},
		r: 1,
	}
	fmt.Println(c.area())

	c = &Circle {
		Point: Point {0, 0},
		r: 0.1,
	}
	fmt.Println(c.area())
	

	// CUBE OUTPUT
	k := &Cube {
		n: 3,
	}
	fmt.Println(k.area())

	k = &Cube {
		n: 100,
	}
	fmt.Println(k.area())

}
