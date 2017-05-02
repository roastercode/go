package main

// Swap example

import (
	"fmt"
)

func swap(x, y *int){  // *int is a part of the type, also says pointer to an int
	
	//tmp := *x
	//*x = *y
	//*y = tmp
        // same as
	*x, *y = *y, *x // *x is an operator

//	xs := []int{} // This is part of the type
//	xs[2]         // This is an operator
}

func rotate(args ...*int) {
	if len(args) == 0 {
		return
	}
	firstValue := *args[0]
	for i := 0; i < len(args)-1; i++ {
		*args[i] = *args[i+1]
	}
	*args[len(args)-1] = firstValue
}

func main(){
	x := 1
	y := 2
	z := 3
	swap(&x, &y)        // swap == rotate
	rotate(&x, &y, &z)  // same
	fmt.Println(x, y, z)
}
