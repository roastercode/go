package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

const (
	miTokm  = 1.60934
)

func main() {
	
	// load from the terminal after a go build ; go install
	// slice 2mi km

	from := os.Args[1]
	to := os.Args[2]

	switch {
	case strings.HasSuffix(from, "mi"):

		switch to {
		case "km":
			miles, _ := strconv.ParseFloat(from[:len(from)-2], 64)
			fmt.Println(miles * miTokm)
		}
	}
	// fmt.Println(from, to)
}
