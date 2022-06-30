package main

import (
	"math"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for i := range ret {
		ret[i] = make([]uint8, dx)
		for j := range ret[i] {
			//ret[i][j] = uint8((i + j) / 2)
			//ret[i][j] = uint8(i * j)
			//ret[i][j] = uint8(i ^ j)
			ret[i][j] = uint8(i * int(math.Log(float64(j))))
			//ret[i][j] = uint8(i % (j + 1))
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
