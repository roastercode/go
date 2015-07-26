package main

import "fmt"

const Tau = 6.28

func main() {
	const World = "World"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Tau, "Day")
	fmt.Print("What is ", Tau, "? The day of Tau\n")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
