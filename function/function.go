package main

import "fmt"

// defining average function
func average(ys []float64) float64 {
	total := 0.0
	for  _, v := range ys {
		total += v
	}
	return total / float64(len(ys))
}

// another function (same as average one) for test
func bubble(sy []float64) (n float64) {
	total := 0.0
	for  _, v := range sy {
		total += v
	}
	n = total / float64(len(sy))
	return
}

func bobble(yy ...float64) (a float64) {
	total := 0.0
	for  _, v := range yy {
		total += v
	}
	a = total / float64(len(yy))
	return
}

// sum function
func sum(bb []float64) float64 {
	total := 0.0
	for _, v := range bb {
		total += v
	}
	return total
}

// defining the function of the table
// calculate by the average function define before

func bilbo() float64 {
	xz := []float64{8, 9, 7, 2, 3}
	o := bubble(xz)
	fmt.Println(o)
	return 0
}

func balbo() float64 {
	u := bobble(18, 9, 7, 2, 3)
	fmt.Println(u)
	return 0
}

func bilba() {
	xs := []float64{98, 93, 77, 82, 83}
	n := bubble(xs)
	fmt.Println(n)
	return
}

func bilbe() float64 {
	zx := []float64{1, 9, 7, 2, 3}
	m := average(zx)
	fmt.Println(m)
	return 0
}

func bilbx() {
	sx := []float64{8, 93, 77, 82, 83}
	p := average(sx)
	fmt.Println(p)
}

// using sum function
func sumx() {
	nn := sum([]float64{98, 65 ,75 ,12 ,45})
	fmt.Println(nn)
}
			
func main(){
	defer bilbx()
	defer balbo()
	defer bilbe()
	defer bilba()
	defer sumx()
	bilbo()
}

