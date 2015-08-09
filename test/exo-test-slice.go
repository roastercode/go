/*

Exercise: Slices - exercise-slices.go

Implement Pic. It should return a slice of length 'dy', each element
of which is a slice of 'dx' 8-bit unsigned integers.
When you run the program, it will display your picture, interpreting
the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include
 '(x+y)/2 x*y, and x^y'

(You need to use a loop to allocate each []uint8 inside the [][]uint9.)

(Use uint8(intValue) to convert between types.)


*/
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy, dz int) [][][]uint8 {
	pic := make([][][]uint8, dy)
	for i := range pic {
		pic[i] = make([][]uint8, dx)
		for j := range pic {
			pic[j] = make([][]uint8, dz)
			for k := range pic {
				pic[i][j][k] = uint8(((j ^ i ^ k) ^ 10) * (10 ^ (i * j * k)))
			}
		}
	}
	return pic
}
func main() {
	pic.Show(Pic)
}
