/*

Switch with no condition

Switch without a condition is the same as switch true.

This construct can be a clean way to write lon if-then-else chains.

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 6:
		fmt.Println("What the are you doing at this time o_O ?!?")
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	case t.Hour() < 22:
		fmt.Println("Good evening!")
	default:
		fmt.Println("What the hell are you doing at this time O_O ?!?")
	}
}

// Own improved version of switch true on time
