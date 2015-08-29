/*

Simple Programme from GopherCon 2015 to test -race
code bug detector:

aurelien@iceberg:~/go/src/github.com/4ur3l13n/go/test$ go run -race simple.go 
==================
WARNING: DATA RACE
Write by goroutine 5:
  runtime.mapassign1()
      /build/go/src/go-1.5/src/runtime/hashmap.go:411 +0x0
  main.main.func1()
      /home/aurelien/go/src/github.com/4ur3l13n/go/test/simple.go:13 +0x60

Previous write by main goroutine:
  runtime.mapassign1()
      /build/go/src/go-1.5/src/runtime/hashmap.go:411 +0x0
  main.main()
      /home/aurelien/go/src/github.com/4ur3l13n/go/test/simple.go:15 +0xb6

Goroutine 5 (running) created at:
  main.main()
      /home/aurelien/go/src/github.com/4ur3l13n/go/test/simple.go:14 +0x76
==================
Found 1 data race(s)
exit status 66


*/

package main

func main() {
	m := make(map[int]int)
	go func() {
		m[1] = 1
	}()
	m[2] = 2
}
