package main

import "fmt"

func do(i interface{}) {
	// type switch(类型选择)是一种按顺序从几个类型断言中选择分支的结构
	// 类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值）
	// 它们针对给定interface值所存储的值的类型进行比较
	// 类型选择中的声明与类型断言 i.(T) 的语法相同，只是具体类型T被替换成了关键字 type
	// 此选择语句判断interface值 i 保存的值类型是 T 还是 S
	// 在 T 或 S 的情况下，变量 v 会分别按 T 或 S 类型保存 i 拥有的值
	// 在默认（即没有匹配）的情况下，变量 v 与 i 所保存值的interface类型和值相同
	switch v := i.(type) {
	case int:
		// %v: 打印值的默认格式
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		// %q: 单引号围绕的字符字面值，由Go语法安全地转义
		// %v: 打印值的默认格式
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		// %T: 相应值的类型的Go语法表示
		fmt.Printf("I don't know about type %T of the variable %v!\n", v, v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
