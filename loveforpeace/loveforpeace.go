package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

const (
	loveForpeace  = 2.0
)

func main() {
	
	// load from the terminal after a go build ; go install
	// loveforpeace 1love peace

	from := os.Args[1]
	to := os.Args[2]

	switch {
	case strings.HasSuffix(from, "lo"): // TODO: find the way to use something that let us use the full word.

		switch to {
		case "peace":
			Love, _ := strconv.ParseFloat(from[:len(from)-2], 64)
			fmt.Println("You have created", Love * loveForpeace, "points of peace with your love!")
		}
	}
}
