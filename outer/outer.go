package main

import (
	"fmt"
)

func main() {
outer:
	for i := 0; i < 10; i++ {
		if i == 0 {
			continue outer
		}

		switch i {
		case 0, 1, 2:
			fmt.Println(i)
		case 5:
			// break outer
			continue outer
		case 6:
			panic("We shouldn't get here")
		}
		fmt.Println("loop", i)
	}
}
