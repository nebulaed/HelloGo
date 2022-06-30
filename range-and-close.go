package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 发送者可通过close关闭一个channel来表示没有需要发送的值了
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	// 接收者可以通过为接收表达式分配第二个参数来测试channel是否被关闭
	// 若没有值可以接收且channel已被关闭，那么在执行完
	// v, ok := <-ch
	// ok会为false
	//for i, ok := <-c; ok; i, ok = <-c {
	//	fmt.Println(i)
	//}

	// 以下语句等价于上面三行语句
	// 循环 for i := range c 会不断从channel接收值，直到它被关闭
	// 注意：只有发送者才能关闭channel，接收者不能。向一个已经关闭的channel发送数据会引发程序panic
	// 注意：channel与文件不同，通常无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要挂你吧。如终止一个range循环
	for i := range c {
		fmt.Println(i)
	}
}
