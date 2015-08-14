/*

Exercise: Errors - exercise-errors.go

Copy your 'Sqrt' function from the earlier http://tour.golang.org/flowcontrol/8
and modify it to return an 'error' value.

'Sqrt' should return a non-nil error value when given a negative number, as it
doesn't support complex numbers.

Create a new type

    type ErrNegativeSqrt float64

and make it an 'error' by giving it a

    func (e ErrNegativeSqrt) Error() string

method such that

'ErrNegativeSqrt(-2).Error()' returns "cannot Sqrt negative number: -2".

Note: a call to 'fmt.Sprint(e) inside the 'Error' method will send the program
into an infinite loop. You can avoid this by converting 'e' fist:
'fmt.Sprintf(float64(e)). Why?

Change your 'Sqrt' function to return an 'ErrNegativeSqrt' value when given a
negative number.

*/

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt positive number: %g", float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, nil
	}
	return 0, ErrNegativeSqrt(f)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
