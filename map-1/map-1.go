/*
Golang     : Slice v4
License    : GNU GPL v3 or later
Author     : Aurélien DESBRIERES
Mail       : aurelien@hackers.camp
Project    : Golang Experimentations
Created on : June 2017

Write with Emacs-Nox ──────────────┐
Go lang ───────────────────────────┘
*/

package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	// Create a map
	countryCapitalMap = make(map[string]string)

	// Insert key-value pairs in the map
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	// Print map using keys
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country, "is", countryCapitalMap[country])
	}

	// Test if entry is present in the map or not
	capital, ok := countryCapitalMap["United States"]
	// If ok is true, entry is present otherwise entry is absent
	if(ok){
		fmt.Println("Capital of United States is", capital)
	} else {
		fmt.Println("Capital of United States is not present")
	}
}

