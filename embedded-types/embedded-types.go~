package main

// Structure and interfaces

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

type Cube struct {
	u int
}

type Rectangle struct {
	l int
	w int
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r*r
}

func anotherCircleArea(f Circle) float64 {
	return math.Pi * f.r * f.r
}

func cubeArea(u int) int {
	return u * u * u
}

func (g *Circle) area() float64 { // struct of a circle with pointer
//	g.r = 100
	return math.Pi * g.r * g.r // modify the calculus for test (works!)
}

func (h *Cube) space() int {
//	h.u = 10
	return h.u * h.u * h.u
}

func (i *Rectangle) univers() int {
	return i.l * i.w
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))

	var c Circle
	fmt.Println(c.x, c.y, c.r)

	d := Circle {
		x: 0,
		y: 0,
		r:10,
	}
	fmt.Println(d)

	// same as

	e := Circle {0, 0, 10}
	fmt.Println(e)

	f := Circle {
		x: 0,
		r: 10,
	}
	fmt.Println(anotherCircleArea(f))

	k := Cube {
		u: 3,
	}
	fmt.Println(cubeArea(k.u))

	g := Circle {
		x: 0,
		r: 3,
	}
	fmt.Println(g.area())

	g = Circle {          // Creating a new circle with the pointer using =
		x: 10,
		r: 7,
	}
	fmt.Println(g.area())

	h := Cube {
		u: 10,
	}
	fmt.Println(h.space())

	h = Cube {             // To create a new cube just use =
		u: 100,
	}
	fmt.Println(h.space())

	i := Rectangle {
		l: 10,
		w: 20,
	}
	fmt.Println(i.univers())

	// Creating a new Rectangle
	i = Rectangle {
		l: 100,
		w: 200,
	}
	fmt.Println(i.univers())
}
