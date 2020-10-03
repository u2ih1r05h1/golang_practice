package main

import "golang.org/x/tour/pic"

import (
	"image"
	"image/color"
)

type Image struct {
	width, height int
	color         uint8
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{img.color + uint8(x), img.color + uint8(y), 255, 255}
}

func main() {
	m := Image{100, 100, 100}
	pic.ShowImage(m)
}
