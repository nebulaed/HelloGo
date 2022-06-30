package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// tree包中定义了类型：
//type Tree struct {
//	Left *Tree
//	Value int
//	Right *Tree
//}

// Walk 步进 tree t将所有的值从tree发送到channel ch
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

// testWalk
//func main() {
//	ch := make(chan int)
//	go Walk(tree.New(1), ch)
//	for i := 0; i < 10; i++ {
//		fmt.Println(<-ch)
//	}
//}

func main() {
	t1, t2 := tree.New(1), tree.New(2)
	fmt.Println(Same(t1, t2))
}
