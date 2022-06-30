package main

import "fmt"

// type assertion(类型断言)提供了访问interface值底层具体值的方式
// t := i.(T)
// 该语句assertion接口值i保存了具体类型T，并将其底层类型为T的值赋予变量t
// 若i并未保存 T 类型的值，该语句就会触发一个panic
// 为了判断一个interface值是否保存了一个特定的类型，type assertion可返回两个值：其底层值以及一个验证assertion是否成功的bool值
// t, ok := i.(T)
// 若i保存了一个T，那么t将会是其底层值，而ok为true
// 否则，ok将为false，而t将为T类型的零值，这时不会触发panic
// 这个语法和读取一个map时的相同之处
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // 报错(panic)
	fmt.Println(f)
}
