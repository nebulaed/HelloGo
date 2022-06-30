package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// 注意: 在 Error 方法内调用 fmt.Sprint(e) 会让程序陷入死循环
	// 可以通过先转换 e 来避免这个问题：fmt.Sprint(float64(e))。这是为什么？
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt 接受到一个负数时，应当返回一个非 nil 的错误值
// 方法使其拥有error值，通过ErrNegativeSqrt(-2).Error()调用该方法应返回错误说明

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	ret := x / 2
	for diff := (ret*ret - x) / (2 * ret); math.Abs(diff) > 0.001; diff = (ret*ret - x) / (2 * ret) {
		ret -= diff
	}
	return ret, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
