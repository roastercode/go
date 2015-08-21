/*

Ranging Over Strings - From GopherCon 2015

*/

package main

import "fmt"

func main() {
	iLoveNY := "I ♥ NY"

	for _, c := range iLoveNY {
		fmt.Printf("%U, %q\n" , c, c) // unicode, quote literal
	}

	// Outputs:
	// U+0049, 'I'
	// U+0020, ' '
	// U+2665, '♥'
	// U+0020, ' '
	// U+004E, 'N'
	// U+0059, 'Y'
}
