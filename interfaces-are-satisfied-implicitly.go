package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 类型通过实现一个interface的所有方法来实现该interface。既然无需专门显式声明，也就没有“implements”关键字
// 隐式interface从interface的实现中解耦了定义，这样interface的实现可以出现在任何包中，无需提前准备
// 因此，也就无需在每一个实现上增加新的interface名称，这样同时也鼓励了明确的interface定义

// 此方法表示类型T实现了interface I，但我们无需显式声明此事
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
