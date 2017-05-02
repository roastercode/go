// reader of file in go
// aurelien@hackers.camp
// read a file in go

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
)

func main() {
	// Request the user the file he wish to read
	fmt.Print("Which file do you want to read? ")
	var input string
	fmt.Scanln(&input)

	// Print the file on the command line output
	f, err := os.Open(input)
	if err != nil {
		log.Fatalln("my program broken", err.Error())
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("my program broken")
	}

	str := string(bs)
	
	fmt.Println(str)
}
