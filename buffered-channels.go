package main

import "fmt"

// channel可以是带buffer的，将buffer长度作为第二个参数提供给make来初始化一个带buffer的channel
// ch := make(chan int, 100)
// 仅当channel的buffer填满后，向其发送数据时才会阻塞。当buffer为空后，接收方会阻塞
// 修改示例填满buffer，看看会发生什么

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 // 会报fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// ch <- 3 // OK!
}
