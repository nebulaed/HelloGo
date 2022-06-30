package main

import (
	"fmt"
	"image"
)

// image 包定义了 Image interface:
//package image
//
//type Image interface {
//	ColorModel() color.Model
//	Bounds() Rectangle
//	At(x, y int) color.Color
//}
// 注意：Bounds方法的返回值Rectangle实际上是个image.Rectangle，它在image包中声明。
// color.Color和color.Model类型也是interface，但通常因为直接使用预定义的实现image.RGBA和image.RGBAModel而被忽视了。
// 这些接口和类型由image/color包定义

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
