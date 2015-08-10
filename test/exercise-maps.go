/*

Exercise: Maps - exercise-maps.go

Implement 'WordCount'. It should return a map of the counts of each "word"
in the string s. The 'wc.Test' function runs a test suite against the provided
function and prints success or failure.

You might find 'strings.Fields' helpful.
http://golang.org/pkg/strings/#Fields

*/
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	a := strings.Fields(s)
	for _, v := range a {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
