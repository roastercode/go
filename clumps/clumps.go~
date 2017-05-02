package main

// Importing format and string package
import (
	"fmt"
	"strings"
)

// Defining WordCount function
func WordCount(str string) map[string]int {
	counts := map[string]int{}
	for _, word := range strings.Fields(str) {
		counts[word]++
	}
	return counts
}

// Defining WordCounter function
func WordCounter(str string) map[string]int {
	counter := map[string]int{}
	for _, word := range strings.Fields(str) {
		counting := counter[word]
		counting = counting + 1
		counter[word] = counting
	}
	return counter
}


func main() {
	// Call of WordCount function
	str := "test test"
	result := WordCount(str)
	fmt.Println(result)

	// Call of WordCounter function
	string := "test test"
	response := WordCounter(string)
	fmt.Println(response)

	// Other form of response of the calling function
	stringofcat := "I have a cat that love cat and catwoman" // able to count just the word without mixed result
	results := WordCount(stringofcat)
	fmt.Println("Number of cats:", results["cat"])
}
