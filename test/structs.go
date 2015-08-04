package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Botox struct {
	Z int
	W int
}

func main() {
	defer fmt.Println(Vertex{1, 2})
	fmt.Println(Botox{3, 4})
}
