/* For
Go  has only one looping construct, the for loop.

The basic for loop looks as it does in C or Java, except that the ( )
are one (they are not even optional) and the { } are required.
*/

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
/*	test := 2
	for 1 < 2 {
		test *= 2
	}
	fmt.Println(test)

// This part create an infinite loop because of the 1 < 2

*/
}
