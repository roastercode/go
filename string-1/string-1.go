/*
Golang - String I
License    : GNU GPL v3 or later
Author     : Aurélien DESBRIERES
Mail       : aurelien@hackers.camp
Project    : Golang Experimentations
Created on : June 2017

Write with Emacs-Nox ──────────────┐
Go lang ───────────────────────────┘
*/

package main

import "fmt"

func main() {
	var greeting = "Hello world!"

	fmt.Printf("normal string: ")
	fmt.Printf("%s", greeting)
	fmt.Printf("\n")
	fmt.Printf("hex bytes: ")
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x ", greeting[i])
	}

	fmt.Printf("\n")

	const sampleText = "\xbd\xb2\x3b\xbc\x20\xe2\x8c\x98"
	/*q flag escape unprintable characters, with + flag it escapes non-ascii characteres as well
          to make ouput unambigous */
	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", sampleText)
	fmt.Printf("\n")
}

