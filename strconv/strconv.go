package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseInt("AF0875", 16, 64)) // hexadecimal interpretation
	fmt.Println(strconv.ParseInt("AF0875x", 16, 64)) // able to see error because of wrong hexadecimal
	fmt.Println(strconv.FormatInt(123, 16)) // interpretation from decimal number to hexadecimal
	fmt.Println(strconv.ParseInt(strconv.FormatInt(123, 16), 16, 64)) // Interpretation & reinterpretation
	// Correctly written of the Interpretation & reinterpretation
	fmt.Println(
		strconv.ParseInt(
			strconv.FormatInt(123, 16),
			16,
			64,
		),
	)
}
