package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 可以为指针接收者声明方法
// 这意味着对于某类型 T，接收者的类型可以用 *T 的文法 （此外，T 不能是像 *int 这样的指针。）
// 指针接收者的方法可以修改接收者指向的值（就像 Scale 在这做的）。由于方法经常需要修改它的接收者，指针接收者比值接收者更常用
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3,4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
