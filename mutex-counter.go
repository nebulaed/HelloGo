package main

import (
	"fmt"
	"sync"
	"time"
)

// 我们已经看到channel非常适合在各个goroutine间进行通信
// 但若我们不需要通信呢？比如说，我们只是想保证每次只有一个goroutine能够访问一个共享变量，从而避免冲突
// 这里涉及的概念叫做互斥(mutual exclusion)，我们通常使用互斥锁(Mutex)这一数据结构来提供这种机制
// Go标准库中提供了sync.Mutex互斥锁类型及其两个方法：
// Lock, Unlock

// SafeCounter的并发使用是安全的
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// 我们可以通过在代码前调用Lock方法，在代码后调用Unlock方法来保证一段代码的互斥执行
// Inc增加给定key的计数器的值
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock之后同一时刻只有一个goroutine能访问c.v
	c.v[key]++
	c.mux.Unlock()
}

// 我们也可以用defer语句来保证互斥锁一定会被解锁
// Value返回给定key的计数器的当前值
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock之后同一时刻只有一个goroutine能访问c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	// 这里用time.Second, time.Millisecond, time.Microsecond, time.Nanosecond效果一样, time.Nanosecond偶尔会得到结果为998，time.Microsecond偶尔会得到999
	// 但对exercise-web-crawler不是，可能因为exercise-web-crawler涉及递归
	time.Sleep(time.Millisecond)
	fmt.Println(c.Value("somekey"))
}
