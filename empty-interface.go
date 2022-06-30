package main

import "fmt"

func main() {
	// 指定了零个方法的interface值被称为空interface
	var i interface{}
	describe(i)

	// 空interface可保存任何类型的值，因为每个类型都至少实现了零个方法
	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

// 空interface被用来处理未知类型的值。
func describe(i interface{}) {
	// 例如，fmt.Printf 可接受类型为 interface{} 的任意数量的参数
	fmt.Printf("(%v, %T)\n", i, i)
}
