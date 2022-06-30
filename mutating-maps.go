package main

import "fmt"

// 修改map
func main() {
	m := make(map[string]int)

	// 在map中插入或修改元素
	m["Answer"] = 42
	// 获取元素
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// 删除元素
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// 通过双赋值检测某个键是否存在
	// elem, ok = m[key]
	// 若 key 在 m 中，ok 为 true ；否则，ok 为 false
	// 若 key 不在map中，那么 elem 是该map元素类型的零值
	// 同样的，当从map中读取某个不存在的键时，结果是map的元素类型的零值
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
