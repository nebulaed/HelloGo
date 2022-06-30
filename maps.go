package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map将键映射到值
// map的零值为nil，nil map既没有键也没有添加键
// make函数会返回给定类型的map，并将其初始化备用
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
