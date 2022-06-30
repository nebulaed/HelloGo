package main

import "fmt"

// slice的零值是 nil
// nil slice的长度和容量为0，且没有底层数组
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
