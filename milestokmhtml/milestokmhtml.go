package main

import (
	"fmt"
)

const (
	miTokm  = 1.60934
)

func main() {
	var number float64
	fmt.Scanln(&number)
	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println(" <head></head>")
	fmt.Println(" <body>")
	// get the number of miles from the user
	fmt.Printf("Miles: %f<br>\n", number)
	fmt.Printf("Kilometers: %f<br>\n", number * miTokm)
	fmt.Println(" </body>")
	fmt.Println("</html>")
}

// to link that code to a html page do:
// install the file.go first
// go build ; go install ; file
// file > /tmp/file.html
