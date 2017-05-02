package main

/* Software that scan your entry, read the file and past the output of the file
   in the terminal.
   Works with the right name of the file in the folder */

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
)

func main() {
	// Scan string
	var file string
	fmt.Printf("Enter the name of file you want to read please: ")
	fmt.Scanln(&file)
	
	f, err := os.Open(file) // Open the file
	if err != nil {
		log.Fatalln("my program broken")
	}
	defer f.Close() // Always close things open

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("my program broken")
	}

	fmt.Printf("%s", bs) // %s convert directly in string the result
}
