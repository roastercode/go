/*

Dmitry Vyukov - Go dynamic Tools - Go fuzz - dv-fuzz.go
GoferCon 2015

*/

package png

import (
	"bytes"
	"image/png"
)

func Fuzz(data []byte) int {
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		if img != nil {
			panic("img != nil on error")
		}
		return 0
	}
	var w bytes.Buffer
	err = png.Encode(&w, img)
	if err != nil {
		panic(err)
	}
	return 1
}
