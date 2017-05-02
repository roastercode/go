// println is a built-in function that does not need to import the library
// the other point is that it could present trouble on remote functionnality
// the point is that it is more clean to use it, but not sure it will stay in the language.
package main

import (
//"fmt"
)

func main() {
	i := 1
	o := 1
	for i <= 10 {
		println(i)
		i = i + 1
	}
	for o <= 10 {
		print("\t", o)
		o = o + 1
	}
	println()

	// Add a variable that is the number
	cnt := 20
	for i <= cnt {
		println(i)
		i = i + 1
	}
	print("\t", i)
	i = i + i
	println()

	ntc := 12
	for m := 0; m <= ntc; m++ {
		print("\t", m)
	}
	println()

	// range is an iteration (répétition)
	// by that way we got the index 'p' and the character 'q'
	for p, q := range "répétition" {
		println(p, q)
	}

	str := "iterate"
	for n := range str {
		print("\t", str[n])
	}
	println()

	cond := true
	for cond {
		println("Hi")
		cond = false
	}

	for {
		println("Hoy!")
		break // without break condition it should loop for ever!
	}
	for {
		continue
		println("Hay!")
		// this one just not stop nor start but eat cpu
	}
}
