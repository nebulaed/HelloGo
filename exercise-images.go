package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Pix    []uint8
	Stride int
	Rect   image.Rectangle
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return i.Rect
}

func (i *Image) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(i.Rect)) {
		return color.RGBA{}
	}
	idx := i.PixOffset(x, y)
	return color.RGBA{i.Pix[idx+0], i.Pix[idx+1], i.Pix[idx+2], i.Pix[idx+3]}
}

func (i *Image) PixOffset(x, y int) int {
	return (y-i.Rect.Min.Y)*i.Stride + (x-i.Rect.Min.X)*4
}

func (i *Image) SetRGBA(x, y int, c color.RGBA) {
	if (!image.Point{x, y}.In(i.Rect)) {
		return
	}
	idx := i.PixOffset(x, y)
	i.Pix[idx+0], i.Pix[idx+1], i.Pix[idx+2], i.Pix[idx+3] = c.R, c.G, c.B, c.A
}

func main() {
	h, w := 256, 256
	m := &Image{
		Pix:    make([]uint8, h*w*4),
		Stride: w * 4,
		Rect:   image.Rect(0, 0, w, h),
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			magnitude := uint8((i + j) / 2)
			//magnitude := uint8(i * j)
			//magnitude := uint8(i ^ j)
			//magnitude := uint8(i * int(math.Log(float64(j))))
			//magnitude := uint8(i % (j + 1))
			m.SetRGBA(i, j, color.RGBA{magnitude, magnitude, 255, 255})
		}
	}
	pic.ShowImage(m)
}
