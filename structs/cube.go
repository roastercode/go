package main

import (
	"fmt"
	"math"
)

type Cube struct {
	u int
}

//func (h *Cube) space() int {
//	return h.u * h.u * h.u
//}

func (h *Cube) space() float64 { // trying to use math.Exp or math.Pow
	return math.Pow(h.u)     // need fix!
}

func main() {
	h := Cube {
		u: 10,
	}
	fmt.Println(h.space())

	h = Cube {            // To create a new cube just use =
		u: 100,
	}
	fmt.Println(h.space())

	h = Cube {
		u: 500,
	}
	fmt.Println(h.space())
}
