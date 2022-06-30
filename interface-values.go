package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	i = &T{"Hello"}
	// (&{Hello}, *main.T)
	// interface也是值。它们可以像其它值一样传递
	// interface值可以用作函数的参数或返回值
	// 在内部，interface值可以看做包含值和具体类型的元组(value, type)
	// interface值保存了一个具体底层类型的具体值
	describe(i)
	// interface值调用方法时会执行其底层类型的同名方法
	i.M()

	i = F(math.Pi)
	// (3.141592653589793, main.F)
	describe(i)
	i.M()
}

func describe(i I) {
	// %T: 相应值的类型的Go语法表示
	fmt.Printf("(%v, %T)\n", i, i)
}
