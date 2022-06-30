package main

import (
	"fmt"
	"time"
)

// 当select中其他分支都没有准备好时，default分支就会执行
// 为在尝试发送或接收时不发生阻塞，可使用default分支:
//select {
//case i := <-c:
//	// 使用i
//default:
//	// 从c中接收会阻塞时执行
//}

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
