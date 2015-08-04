/*

Stacking defers

Deferred function calls are pushed onto a stack.
When a function returns, its deferred calls are executed
in las-in-first-out order.

To learn more about defer statements read this
http://blog.golang.org/defer-panic-and-recover

*/

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
