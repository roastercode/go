package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(strings.ToTitle("Hello World"))
	fmt.Println(strings.ToLower("hello world"))
	fmt.Println(strings.Replace("Hello World Hello World Hello World", "Hello", "Come", 2))
	fmt.Println(strings.Replace("    Aero    space", "    ", "", 2)) // Remove space
	fmt.Println(strings.TrimSpace("    Aero    space")) // Also remove space but in the front
}
