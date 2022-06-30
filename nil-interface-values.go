package main

import "fmt"

type I interface {
	M()
}

func main() {
	// nil interface值
	var i I
	// 函数实参是值为nil的指针并不会出错
	describe(i)
	// 为nil interface调用方法会产生runtime错误，因为interface的元组内并未包含能够指明该调用哪个具体方法的类型
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
