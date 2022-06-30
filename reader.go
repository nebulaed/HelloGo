package main

import (
	"fmt"
	"io"
	"strings"
)

// io包指定了io.Reader interface，它表示从数据流的末尾进行读取
// Go标准库包含了该interface的许多实现，包括文件、网络连接、压缩和加密等
// io.Reader interface有一个Read方法
//func (T) Read(b []byte) (n int, err error)
// Read用数据填充给定的字节slice并返回填充的字节数和错误值
// 在遇到数据流的末尾时，即本次连1个字符也读不到时，它会返回一个io.EOF错误
// 示例代码创建了一个strings.Reader并以每次8字节的速度读取它的输出
func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		// 字符字面值，由Go语法安全地转义，如0x4E2D输出为"中"
		// 这里不能打印b，否则最后一次没读到东西b也会有内容
		// 每次Read只会写入到b[0:n]中
		fmt.Printf("b[:n] = %q\n", b[:n])
		fmt.Printf("b = %q\n", b)
		if err == io.EOF {
			break
		}
	}
}
