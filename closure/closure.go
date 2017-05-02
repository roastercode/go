package main

import "fmt"

// Increasing Number Generator
func increasingNumbersGenerator() func() int {
	y := 0
	return func() int {
		y++
		return y
	}
}

// generator starting from 0
func makeEvenGenerator() func() uint {
	i := uint(0) // or var i uint
	return func() (ret uint) {
		ret = i
		i += 3 // Increase by 3 each time
		return
	}
}

// generator starting from the number
func makeAnotherGenerator() func() int {
	o := int(0)
	return func() int {
		o += 4
		return o
	}
}

// generator starting from 0
func nextGenerator() func() int {
	k := int(0) // or var i uint
	return func() (ret int) {
		ret = k
		k += 5 // Increase by 5 each time
		return
	}
}

// multiplicator generator starting from the number
func multigenerator() func() int {
	b := int(8)
	return func() int {
		b *= 2
		return b
	}
}

// dividor generator
func dividergenerator() func() float64 {
	c := float64(2)
	return func() float64 {
		c /= float64(3)
		return c
	}
}

func filter(numbers []int, callback func(int) bool ) []int {
	xs := []int{}
	for _, ys := range numbers {
		if callback(ys) {
			xs = append(xs, ys)
		}
	}
	return xs
}



// Function including different function + the generator
func main(){
	// Included function in the function
	half := func(n int) (int, bool) {
		return n/2, n%2 == 0
	}
	fmt.Println(half(2))
	{
		// Create another function in the function
		x := 0
		increment := func() int {
			x++
			return x
		}
		fmt.Println(increment())
		fmt.Println(increment())
		fmt.Println(increment())
		fmt.Println(increment())
		fmt.Println(increment())
	}
	
	// Call of the 1st generator function
	numbers := increasingNumbersGenerator()
	fmt.Printf("%T\n", numbers) // requesting what we are using
	fmt.Println(numbers())
	fmt.Println(numbers())

	// Call of the 2nd generator function
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 3
	fmt.Println(nextEven()) // 6

	// Call of the 3rd generator function
	oGenerator := makeAnotherGenerator()
	fmt.Println(oGenerator()) // 4
	fmt.Println(oGenerator()) // 8
	fmt.Println(oGenerator()) // 12

	// Call of the 4th generator function
	anotherGenerator := nextGenerator()
	fmt.Println(anotherGenerator()) // 0
	fmt.Println(anotherGenerator()) // 5
	fmt.Println(anotherGenerator()) // 10

	// Call of the 5th generator function
	multiplicator := multigenerator()
	fmt.Println(multiplicator()) // 16
	fmt.Println(multiplicator()) // 32
	fmt.Println(multiplicator()) // 64

	// Call of the 6th generator function
	dividor := dividergenerator()
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())
	fmt.Println(dividor())

	// calling callback function
	max := 1
	xs := filter([]int{1, 2, 3, 4}, func(ys int) bool {
		return ys > max
	})
	fmt.Println(xs)
}
