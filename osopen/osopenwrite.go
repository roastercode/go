package main

// NEED TO BE FIXED
// Trying to use WriteFile (from ioutil) to write data in a file

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
)

func main() {
	f, err := os.Open("hello.txt") // Open the file
	if err != nil {
		log.Fatalln("my program broken")
	}
	defer f.Close() // Always close things open

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("my program broken")
	}

	fmt.Printf("%s", bs)

	bw, err := ioutil.WriteFile(f), data[Hey]byte, perm os.FileMode
	if err != nil {
		log.Fatalln("my program don't write")
	}

}
