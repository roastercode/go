package main

import (
	"fmt"
)

func main() {
	println("Hello Bob!")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			println(i, "even")
		} else if i%3 == 0 {

		} else {
			println(i, "odd")
		}
	}
	println()

	for o := 1; o <= 10; o++ {
		if o%2 == 0 {
			println(o, "even")
		}
	}
	println()

	for e := 1; e <= 10; e++ {
		if e == 0 {
			println("Zero")
		} else if e == 1 {
			println("One")
		} else if e == 2 {
			println("Two")
		} else if e == 3 {
			println("Three")
		} else if e == 4 {
			println("Four")
		} else if e == 5 {
			println("Five")
		} else if e == 6 {
			println("×")
		} else if e == 7 {
			println("œ")
		} else if e == 8 {
			println("¥")
		} else if e == 9 {
			println("Nine!")
		} else if e == 10 {
			println("You are at the fucking limit boy!")
		} else if e == 11 { // This line should not print since you are out of range!
			println("How the hell is that possible o_O")
		}
	}
	println()

	// Now same with swtich
	for a := 1; a <= 10; a++ {
		switch a {
		case 0:
			println("Zero")
			fallthrough
		case 1:
			println("One")
		case 2:
			println("Two")
		case 3:
			println("Three")
		case 9:
			println("Nine")
		default:
			println("Unknown Number")
		}
	}
	b := 10
	switch {
	case b == 42:
	case b == 43:
	case b == 54:
	case b > 100:
	}
	fmt.Println("Hi")
}
