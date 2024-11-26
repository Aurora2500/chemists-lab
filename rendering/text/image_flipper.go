package text

import (
	"image"
	"image/color"
	"image/draw"
)

type FlippedImage struct {
	Image draw.Image
}

func (fi *FlippedImage) ColorModel() color.Model {
	return fi.Image.ColorModel()
}

func (fi *FlippedImage) Bounds() image.Rectangle {
	return fi.Image.Bounds()
}

func (fi *FlippedImage) At(x, y int) color.Color {
	b := fi.Image.Bounds()
	newY := b.Min.Y + b.Max.Y - 1 - y
	return fi.Image.At(x, newY)
}

func (fi *FlippedImage) Set(x, y int, c color.Color) {
	b := fi.Image.Bounds()
	newY := b.Min.Y + b.Max.Y - 1 - y
	fi.Image.Set(x, newY, c)
}
