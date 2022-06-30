package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// goroutine是由Go Runtime管理的轻量级线程——协程
// go f(x, y, z)
// 会启动一个新的goroutine并执行f(x, y, z)
// f, x, y, z的求值发生在当前的goroutine中，而f的执行发生在新的goroutine中

// goroutine在相同的地址空间中运行，因此在访问共享内存时必须进行同步。
// sync包提供了这种能力，不过在Go中并不经常用到，因为还有其他办法

func main() {
	go say("world")
	say("Hello")
}
