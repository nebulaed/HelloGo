package main

import (
	"fmt"
	"math"
)

// interface类型 是由一组方法签名定义的集合
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// interface类型的变量可以保存任何实现了这些方法的值
	a = f // a MyFloat实现了Abser
	fmt.Println(a.Abs())
	a = &v // a *Vertex实现了Abser

	// 下面一行，v是一个Vertex(而不是*Vertex)，没有实现Abser
	// 由于 Abs 方法只为 *Vertex （指针类型）定义，因此 Vertex（值类型）并未实现 Abser
	//a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
