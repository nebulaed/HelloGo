package main

import "fmt"

// slice拥有 长度 和 容量
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取slice使其长度为0
	// slice的长度就是其包含的元素个数
	// slice的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数
	// slice的长度和容量可通过表达式len(s)和cap(s)来获取
	// 你可以通过重新切片来扩展一个slice，给它提供足够的容量
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)
}
