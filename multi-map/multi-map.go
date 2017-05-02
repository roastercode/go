package main

import (
	"fmt"
)

func main() {

	// INTRODUCTION
	fmt.Println("*\tMULTI MAP WITH THE PERIODIC TABLE\t*")


	// DEFINITION OF THE MAP WITH INT
	// a:= map[int]int{
	// 	7: 10,
	//	3: 0,
	// }


	// ORIGINAL MULTI-MAP WITH STRING ONLY
	// elements := map[string]map[string]string{
	//  	"H": map[string]string{
	//		"name":"Hydrogen",
	//		"state":"gas",
	//	},
	

	// MODIFIED FROM THE ORIGINAL WITH AN INT MORE
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":"Hydrogen",
			"state":"gas",
//			"atomic number":"1",
		},
		"He": map[string]string{
			"name":"Helium",
			"state":"gas",
//			"atomie number":"2",
		},
		"Li": map[string]string{
			"name":"Lithium",
			"state":"solid",
//			"atomic number":"3",
		},
		"Be": map[string]string{
			"name":"Beryllium",
			"state":"solid",
//			"atomic number":"4",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	// ADDING SCANF FUNCTION
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
	//fmt.Println(elements["*"]) // Print an empty line
	//fmt.Println(elements[""])  // Print an empty line

	// TEST ON MULTIPLY
	// fmt.Println(elements["H"]["He"]) create error as it or separate with ','
	//fmt.Println(elements{["Hi"] ["He"]}) //create error as unexpected comma and / or brackets
}
