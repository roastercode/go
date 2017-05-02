package main

import (
	"fmt"
)

func main() {

	// INTRODUCTION
	fmt.Println("*\tMAPS IS A DICTIONNARY\t*")
	fmt.Println("*\tCLEANER ARRAY SOLUTION\t*")

	// DEFINITION OF THE MAP
	a:= map[int]int{
		7: 10,
		3: 0,
	}

	// PRINTING THE MAP AS IT IS
	fmt.Println(a)
	
	fmt.Println("*\tMOVE THINGS USING DELETE\t*")
	
	// CALL OF DELETE BUILTIN FUNCTION
	delete(a, 7)
	fmt.Println(a)


	// MAP DELETE - YET ANOTHER EXAMPLE
	fmt.Println("*\tYET ANOTHER EXAMPLE OF MAP DELETE\t*")

	// DEFINITION OF THE VARIABLE ELEMENT
	elements := make(map[string]string)

	// MAP
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	// CALL TO PRINT AN ELEMENT
	fmt.Println(elements["Li"])

	// ADDING SCANF FUNCTION
	// &elements := [abbr]
	fmt.Printf("Write a physical element abreviation please: ")
	// ADDING THE SYMBOL STRING FOR THE SCANF FUNCTION
	var symbol string

	// CALL OF THE SYMBOL
	n, err := fmt.Scanf("%v", &symbol)

	// RETURN ERROR STATUS FAILED TO SCAN ANY VALUES
	if err != nil {
		fmt.Println("failed to scan any correct symbol", n, err, symbol)
	}

	// RETURN ERROR UNKNOWN ELEMENT
	if _, ok := elements[symbol]; !ok {
		fmt.Println("unknown element symbol")
	} else {
		fmt.Println("Here is the the element you request for:", elements[symbol])
		// RETURN THE RIGHT ANSWER
	}
	
	// TEST ON EMPTY []
	fmt.Println(elements["*"]) // Print an empty line
	fmt.Println(elements[""])  // Print an empty line

	// TEST ON MULTIPLY
	// fmt.Println(elements["H"]["He"]) create error as it or separate with ','
	//fmt.Println(elements{["Hi"] ["He"]}) //create error as unexpected comma and / or brackets
}
