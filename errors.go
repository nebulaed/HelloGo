package main

import (
	"fmt"
	"time"
)

// Go 程序使用 error 值来表示错误状态
// 与 fmt.Stringer 类似，error 类型是一个内建interface
//type error interface {
//	Error() string
//}
// 与 fmt.Stringer 类似，fmt 包在打印值时也会满足 error

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// 这里返回类型是error，所以进行了一次隐式赋值
// var e error = &MyError{time. Now(), "it didn't work"}
// 这样的好处在于，解耦了interface error的实现，其实际实现error.Error()
// 由保存的值的类型定义的T.Error()决定
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// 通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 nil 来进行错误处理
	// error 为 nil 时表示成功；非 nil 的 error 表示失败
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
