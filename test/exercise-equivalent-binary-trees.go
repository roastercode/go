/*

Exercise: Equivalent Binary Trees

1. Implement the 'Walk' function.

2. Test the 'Walk' funtion.

The function 'tree.New(k)' constructs a randomly-structured binary tree holding
the values k. '2k, 3k, ..., 10k.

Create a new channel 'ch' and kick off the walker:

    go Walk(tree.New(1), ch)

Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

3. Implement the 'Same' function using 'Walk' to determine whether 't1' and 't2' store the
same values.

4. Test the 'Same' function.

'Same(tree.New(1), tree.New(1)) should return true, and 'Same(tree.New(1), tree.New(2))
should return false.

*/

package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)

}

func _walk(t *tree.Tree, ch chan int) {
	if t != nil {
		_walk(t.Left, ch)
		ch <- t.Value
		_walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <- ch2 {
			return false
		}
	}
	return true
}

func main() {
	//tree.New(2)
	ch := make(chan int)	
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Print(v)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
