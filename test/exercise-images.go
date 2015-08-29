/*

Exercise: Image - exercise-image.org

Remember the picture generator you wrote earlier? Let's write another one,
but this time it will return an implementation of 'image.Image' instead of a
slice of data.



*/

package main

import (
	"image"
	"golang.org/x/tour/pic"
	"image/color"
)

type Image struct{
	Width, Height int
	colr uint8
}

func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}

func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (r *Image) At(x, y int) color.Color {
	return color.RGBA{r.colr+uint8(x), r.colr+uint8(y), 255, 255}
}

func main() {
	m := Image{200, 200, 255}
	pic.ShowImage(&m)
}
