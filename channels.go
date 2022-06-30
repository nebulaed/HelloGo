package main

import "fmt"

// channel是带有类型的管道，可以通过它用信道操作符<-来发送或接收值
// ch <- v	//将v发送至channel ch
// v := <-ch //从ch接收值并赋予v
// (箭头就是数据流的方向)
// 和map与slice一样，channel在使用前必须创建：
// ch := make(chan int)
// 默认情况下，发送和接收操作在另一端准备好之前都会阻塞
// 这使得goroutine可以在没有显式的锁或条件变量的情况下进行同步

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将sum送入c这个channel
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	// 这里将求和任务分配给两个goroutine，一旦两个goroutine完成了它们的计算就能得到最终结果
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
