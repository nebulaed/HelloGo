package main

import "fmt"

// 返回一个"返回int的函数"
func fibonacci() func(int) int {
	prev := 0
	num := 0
	return func(x int) int {
		temp := prev
		prev = num
		if x == 1 {
			num = 1
		} else {
			num += temp
		}
		return num
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
