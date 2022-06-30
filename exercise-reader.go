package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// 给MyReader添加一个Read([]byte) (int, error) 方法
func (rd MyReader) Read(b []byte) (int, error) {
	for id := range b {
		b[id] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
