package main

import "fmt"

// slice就像array的引用
// slice并不存储任何数据，它只是描述了底层array中的一段
// 与它共享底层array的slice都会观测到这些修改
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
