package main

import "fmt"

// 为slice追加新的元素是种常用的操作，为此 Go 提供了内建的 append 函数
// append 的第一个参数 s 是一个元素类型为 T 的slice，其余类型为 T 的值将会追加到该slice的末尾
// append 的结果是一个包含原slice所有元素加上新添加元素的slice
// 当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的slice会指向这个新分配的数组
func main() {
	var s []int
	printSlice(s)

	// 添加一个空slice
	s = append(s, 0)
	printSlice(s)

	// 这个slice会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性增加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s)
}
