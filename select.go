package main

import "fmt"

// select语句会使一个goroutine可以等待多个通信操作
// select会阻塞到某个分支可以继续执行为止，这时就会执行该分支
// 当多个分支都准备好时会随机选择一个执行

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// 另起一个协程部分必须在当前协程前面，否则会报fatal error: all goroutines are asleep - deadlock!
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
