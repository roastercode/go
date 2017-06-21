/*
Golang - function v2
License    : GNU GPL v3 or later
Author     : Aurélien DESBRIERES
Mail       : aurelien@hackers.camp
Project    : Golang Experimentations
Created on : June 2017

Write with Emacs-Nox ──────────────┐
Go lang ───────────────────────────┘
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	greetings := []string{"Hello","world!"}
	fmt.Println(strings.Join(greetings, " "))
}
