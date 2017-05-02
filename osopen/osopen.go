package main

// Using ioutil and log to read a file

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"strings"
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

	fmt.Printf("%s", bs) // %s convert directly in string the result
	fmt.Println(bs)      // println print in byte but can be converted in string

	str := string(bs)    // convert to string
	fmt.Println(str)     // print the file as it is

	// Another form of reader to ad to ReadAll from packages strings
	rdr := strings.NewReader("some string to read")

	bo, err := ioutil.ReadAll(rdr)
	if err != nil {
		log.Fatalln("my program broken")
	}

	stg := string(bo)
	fmt.Println(stg)

}
