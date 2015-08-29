/*

Buffered Channels - buffered-channels.go

Channels can be 'buffered'. Provide the buffer length as the second
argument to 'make' to initialize a buffered channel:

    ch := make(chan in, 100)

Sends to a buffered channel block only when the buffer is full. Receiveds
block when the buffer is empty.

Modify the example to overfill the buffer and see what happens.

*/

package main

import "fmt"

func main() {
	ch := make(chan int, 30000000000) // Buffer out of range creating a panic
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
