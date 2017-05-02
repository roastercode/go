package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
)

const (
	loTope  = 2.0
)

func main() {
	number, err := strconv.ParseFloat(os.Args[1], 64) // Enter the number directly in the command
	                                                  // go run lovetopeace.go 45
	if err != nil {
		panic(err); log.Fatalln(err)
	}
	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println(" <head></head>")
	fmt.Println(" <body>")
	// get the number of peace from the user
	fmt.Printf("Love: %f<br>\n", number)
	fmt.Printf("Peace: %f<br>\n", number * loTope)
	fmt.Println(" </body>")
	fmt.Println("</html>")
}

// to link that code to a html page do:
// install the file.go first
// go build ; go install ; file
// file > /tmp/file.html
