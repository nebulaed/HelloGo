package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s)

	// 在进行切片时，你可以利用它的默认行为来忽略上下界
	// slice下界的默认值为 0，上界则是该slice的长度。
	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
