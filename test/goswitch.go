/*

Switch

You probably knew what switch was going to look like.

A case body breaks automatically, unless it ends with a fallthrough statement.

*/

package main

import (
	"fmt"
	"runtime"
)

// Introduction to import runtime

func main() {
	fmt.Print("Go run on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("GNU / Linux.")
	case "freebsd":
		fmt.Println("FreeBSD.")
	case "openbsd":
		fmt.Println("OpenBSD.")
	case "plan9":
		fmt.Println("Plan 9.")
	case "windows":
		fmt.Println("Holly Shit.")
	default:
		// Add other type of OS
		// freebsd, openbsd
		// plan9, windows
		fmt.Printf("%s.", os)
	}
}
