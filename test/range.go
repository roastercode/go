/*

Range - range.go

The 'range' form of the 'for' loop iterates over a slice or map.

*/

package main

import "fmt"

var pew = []int{1, 2, 4, 8, 16, 32, 64, 128}
var paw = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pew {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	for i, v := range paw {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
