package main

/*
Software that scan your entry, read the file and past the output of the file
in the terminal.
*/

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
)

func main() {

/*

gocat filename
os.Args[0] == gocat (so we do not use 0 that call for the name of the program itself
os.Args[1] == filename (1 call for the filename in every sort of directory you tell

*/

	f, err := os.Open(os.Args[1]) // Open the file
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
