/*

Exercise: rot13Reader - exercise-rot-reader.go

A common pattern is an 'io.Reader' http://golang.org/pkg/io/#Reader
that wraps another 'io.Reader', modifying the stream in some way.

For example, the 'gzip.NewReader' http://golang.org/pkg/compress/gzip/#NewReader
 function takes an 'ioReader' (a stream of compressed data) and returns a
'*gzip.Reader' that also implements 'io.Reader' (a stream of the decompressed
data).

Implement a 'rot13Reader' that implement 'io.Reader' and reads from an
'io.Reader', modifying the stream by applying the 'rot13' 
http://en.wikipedia.org/wiki/ROT13 substitution cipher to all alphabetical
characters.

The 'rot13Reader' type is provided for you. Make it an 'io.Reader' by implementing
its 'Read' method.

*/

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Coder struct {
	r io.Reader
}

func (rot *rot13Coder) Read(p []byte) (n int, err error) {
	n,err = rot.r.Read(p)
	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] < 'N') || (p[i] >='a' && p[i] < 'n') {
			p[i] += 13
		} else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z'){
			p[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader(
		"Encode the crack!\n")
	r := rot13Coder{s}
	io.Copy(os.Stdout, &r)
}
