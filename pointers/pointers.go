package main

// Simple pointer

import (
	"fmt"
)

func zero(xPtr *int) { // The * create the pointer
	*xPtr = 1
}

func bubble(yBubble *float64) {
	*yBubble = 10.53
}

func fillWithOne(b [100]byte) {  // function to request to print 100 times
	fmt.Printf("adress in fill with one: %p\n", &b) // %p for pointer (cf. godoc/pkg/fmt)
	for i := range b {
		b[i] = 1
	}
}

// function of a pointer that request 100 times the same byte 
func pointToTheSame(c *[100]byte) {
	fmt.Printf("address in fill with one: %p\n", c)
	for d := range *c {
		(*c)[d] = 1
	}
}

func WithOne(f []byte) {
	fmt.Printf("address in fill with one: %p\n", &f)
	for h := range f {
		f[h] = 1
	}
}

func square(k *float64) {
	//	*k = float64(*k) * float64(*k) // works but ↓ is simpler
	*k = *k * *k
}


func main() {
	x := 5
	zero(&x)
	fmt.Println(x)

	y := 2.12
	bubble(&y)
	fmt.Println(y)

	z := 5
	fmt.Println(z)
	fmt.Println(&z)  // referencing
	fmt.Println(*&z) // deferencing

	//a := [100]byte{} // Printing 100 times
	a := &[100]byte{} // Creating a pointer &
	//fmt.Printf("address in main: %p\n", &a)
	fmt.Printf("address in main: %p\n", *a) // Calling to print the pointer
	//fillWithOfne(a)
	//fmt.Println(a, &a)

	e := [100]byte{}
	fmt.Printf("address in main: %p\n", &e)
	pointToTheSame(&e) // Calling the function of the pointer
	fmt.Println(e)

/*	var j int
	jPtr := new([100]byte) // => return a pointer of 100 bytes
	fmt.Println(jPtr)
	*jPtr = j
	jPtr = &j */

	g := make([]byte, 100) // => return a slice of byte
	fmt.Printf("address in main: %p\n", &g)
	WithOne(g)
	fmt.Println(g)

	k := 1.5
	square(&k)
	fmt.Println(k)
}
