package main

import "fmt"

func main() {
	// slice可以用内建函数 make 来创建，这也是你创建动态数组的方式

	// make 函数会分配一个元素为零值的数组并返回一个引用了它的slice
	a := make([]int, 5)
	printSlice("a", a)

	// 要指定它的容量，需向 make 传入第三个参数
	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	// 从前面截会减少cap，从后面截不会
	d := c[2:5]
	printSlice("d", d)

	// 从前面截会减少cap，从后面截不会
	e := d[:2]
	printSlice("e", e)
}

func printSlice(name string, s []int) {
	fmt.Printf("%s, len=%d, cap=%d, %v\n", name, len(s), cap(s), s)
}
