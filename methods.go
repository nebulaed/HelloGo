package main

import (
	"fmt"
	"math"
)

// Go 没有类。不过你可以为结构体类型定义方法
type Vertex struct {
	X, Y float64
}

// 方法就是一类带特殊的 接收者 参数的函数
// 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间
// 在此例中，Abs 方法拥有一个名为 v，类型为 Vertex 的接收者
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func Abs_func(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := Vertex{3, 4}
	// 普通函数
	fmt.Println(Abs_func(v))
	// 结构体的方法
	fmt.Println(v.Abs())
}
