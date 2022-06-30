package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map的文法与struct相似，不过必须有键名
var m = map[string]Vertex {
	"Bell Labs": Vertex{
		40.64833, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
